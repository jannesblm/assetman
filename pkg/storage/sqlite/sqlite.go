package sqlite

import (
	"errors"
	"github.com/cmp307/assetman/pkg/auth"
	"github.com/cmp307/assetman/pkg/storage"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
)

// Access errors
var (
	ErrNotInitialized    = errors.New("database not initialized")
	ErrNoWritePermission = errors.New("no write permissions")
)

type Repository interface {
	Connect() (*gorm.DB, error)
}

// DB is decorating gorm.DB to add auth and initialisations checks.
type DB struct {
	*gorm.DB
	auth        auth.Service
	initialized bool
}

func (db *DB) MigrateAll() error {
	return db.AutoMigrate(
		storage.Asset{},
		storage.Manufacturer{},
		storage.HardwareAsset{},
		storage.SoftwareAsset{},
		storage.User{},
		storage.Report{},
		storage.Cve{},
	)
}

// Bypass disables authentication checks for the session in tx.
// This should be used with care and not exposed to frontend.
func (db *DB) Bypass(tx func(tx *gorm.DB) error) error {
	return db.WithContext(context.WithValue(context.TODO(), "BypassAuth", true)).Transaction(tx)
}

func (db *DB) SetInitialized(state bool) {
	db.initialized = state
}

func (db *DB) checkInitialized(gorm *gorm.DB) {
	if !db.initialized {
		gorm.AddError(ErrNotInitialized)
	}
}

// checkWritePermissions calls auth.Service.GetUser and checks if the
// currently authenticated User has the "admin" group (allowing write-permissions)
func (db *DB) checkWritePermissions(gorm *gorm.DB) {
	if bypass, ok := gorm.Statement.Context.Value("BypassAuth").(bool); ok && bypass {
		return
	}

	if user, err := db.auth.GetUser(); err == nil && user.Group == "admin" {
		return
	}

	gorm.AddError(ErrNoWritePermission)
}

// InitialiseCallbacks registers initialisation checks before CRUD and
// authorisation checks before CUD operations.
func (db *DB) InitialiseCallbacks() {
	db.Callback().Query().Before("*").Register("check_init", db.checkInitialized)
	db.Callback().Create().Before("*").Register("check_init", db.checkInitialized)
	db.Callback().Update().Before("*").Register("check_init", db.checkInitialized)
	db.Callback().Delete().Before("*").Register("check_init", db.checkInitialized)

	db.Callback().Create().Before("*").Register("check_auth", db.checkWritePermissions)
	db.Callback().Update().Before("*").Register("check_auth", db.checkWritePermissions)
	db.Callback().Delete().Before("*").Register("check_auth", db.checkWritePermissions)
}

func Connect(path string, auth auth.Service) (*DB, error) {
	gorm, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return &DB{}, err
	}

	db := &DB{gorm, auth, false}
	db.InitialiseCallbacks()

	return db, nil
}

type repository struct {
	db *DB
}

type assetRepository struct {
	repository
}

type manufRepository struct {
	repository
}

type userRepository struct {
	repository
}

type reportRepository struct {
	repository
}

func NewAssetRepository(db *DB) *assetRepository {
	return &assetRepository{
		repository{db},
	}
}

func NewManufacturerRepository(db *DB) *manufRepository {
	return &manufRepository{
		repository{db},
	}
}

func NewUserRepository(db *DB) *userRepository {
	return &userRepository{
		repository{db},
	}
}

func NewReportRepository(db *DB) *reportRepository {
	return &reportRepository{
		repository{db},
	}
}

// paginate is a generic pagination method that sets
// Order, Limit and Offset from the storage.QueryOptions struct
func (t repository) paginate(options storage.QueryOptions) *gorm.DB {
	tx := t.db.Preload(clause.Associations)

	if len(options.Order) > 0 {
		tx.Order(options.Order)
	}

	if options.Limit > 0 {
		tx.Limit(options.Limit)
	}

	if options.Offset > 0 {
		tx.Offset(options.Offset)
	}

	return tx
}

func (t *repository) checkWhitelist(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

// AssetRepository

func (t *assetRepository) GetById(id uint) storage.Asset {
	var asset storage.Asset

	t.db.Preload(clause.Associations).
		Preload("HardwareAsset.InstalledSoftware").
		Where("id = ?", id).
		First(&asset)

	return asset
}

func (t *assetRepository) GetAllSoftware() ([]storage.Asset, error) {
	var all []storage.Asset

	err := t.db.Preload("SoftwareAsset").
		Select("ID", "Name", "AssetID").
		Where("asset_type = ?", storage.TypeSoftware).
		Find(&all).
		Error

	return all, err
}

// Paginate paginates assets by type storage.TypeHardware or storage.TypeSoftware
func (t *assetRepository) Paginate(typ string, options storage.QueryOptions) ([]storage.Asset, error) {
	var assets []storage.Asset

	tx := t.paginate(options)

	if typ != storage.TypeHardware && typ != storage.TypeSoftware {
		return []storage.Asset{}, errors.New("invalid type")
	}

	tx.Where("asset_type = ?", typ)

	if len(options.Query) > 0 {
		if len(options.QueryField) == 0 {
			options.QueryField = "name"
		}

		fieldWhitelist := []string{
			"assets.name",
			"type_name",
			"description",
			"purchase_date",
			"Manufacturer.name",
			"HardwareAsset.ip",
			"HardwareAsset.mac",
			"HardwareAsset.model_name",
			"SoftwareAsset.version",
			"SoftwareAsset.license_type",
			"SoftwareAsset.license_key",
		}

		if t.checkWhitelist(fieldWhitelist, options.QueryField) {
			tx.Where(options.QueryField+" like ?", "%"+options.Query+"%")
		}
	}

	err := tx.Joins("Manufacturer").
		Joins("HardwareAsset").
		Joins("SoftwareAsset").
		Preload("Report.Cves").
		Find(&assets).
		Error

	return assets, err
}

func (t *assetRepository) CountSoftware() int64 {
	var count int64

	t.db.Model(&storage.Asset{}).
		Where("asset_type = ?", storage.TypeSoftware).
		Count(&count)

	return count
}

func (t *assetRepository) CountHardware() int64 {
	var count int64

	t.db.Model(&storage.Asset{}).
		Where("asset_type = ?", storage.TypeHardware).
		Count(&count)

	return count
}

func (t *assetRepository) CountAll() int64 {
	var count int64
	t.db.Model(&storage.Asset{}).Count(&count)

	return count
}

// Save uses reflection to save the associated Asset type.
// A new or existing manufacturer will be linked.
// The InstalledSoftware association will be replaced.
func (t *assetRepository) Save(asset storage.Asset) (uint, error) {
	err := t.db.Transaction(func(tx *gorm.DB) error {
		err := t.db.
			Model(reflect.New(asset.Type()).Interface()).
			Clauses(
				clause.OnConflict{
					UpdateAll: true,
				}).
			Create(asset.Polymorphic()).
			Error

		err = t.db.FirstOrCreate(&asset.Manufacturer,
			&storage.Manufacturer{
				Name: asset.Manufacturer.Name,
			},
		).Error

		if asset.AssetType == storage.TypeHardware {
			installed := asset.HardwareAsset.InstalledSoftware

			err = t.db.Model(asset.HardwareAsset).
				Association("InstalledSoftware").
				Clear()

			err = t.db.Model(asset.HardwareAsset).
				Association("InstalledSoftware").
				Append(installed)
		}

		if err != nil {
			return err
		}

		asset.AssetID = asset.Polymorphic().GetID()
		asset.ManufacturerID = asset.Manufacturer.ID

		return t.db.
			Omit("Report", "HardwareAsset", "SoftwareAsset").
			Save(&asset).
			Error
	})

	return asset.ID, err
}

func (t *assetRepository) Delete(asset storage.Asset) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		err := t.db.
			Model(reflect.New(asset.Type()).Interface()).
			Unscoped().
			Delete(asset.Polymorphic()).
			Error

		if err != nil {
			return err
		}

		return t.db.
			Unscoped().
			Delete(&asset).
			Error
	})
}

// ManufacturerRepository

func (t *manufRepository) GetById(id uint) storage.Manufacturer {
	var manufacturer storage.Manufacturer

	t.db.Preload(clause.Associations).
		Where("id = ?", id).
		First(&manufacturer)

	return manufacturer
}

func (t *manufRepository) GetAll() ([]storage.Manufacturer, error) {
	var all []storage.Manufacturer
	err := t.db.Find(&all).Error

	return all, err
}

func (t *manufRepository) Paginate(options storage.QueryOptions) ([]storage.Manufacturer, error) {
	var manufacturers []storage.Manufacturer

	tx := t.paginate(options)

	if len(options.Query) > 0 {
		fieldWhitelist := []string{
			"name",
		}

		if t.checkWhitelist(fieldWhitelist, options.QueryField) {
			tx.Where(options.QueryField+" like ?", "%"+options.Query+"%")
		}
	}

	err := tx.Find(&manufacturers).Error

	return manufacturers, err
}

func (t *manufRepository) CountAll() int64 {
	var count int64
	t.db.Model(&storage.Manufacturer{}).Count(&count)

	return count
}

func (t *manufRepository) Delete(manufacturer storage.Manufacturer) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&storage.Asset{}).
			Where("manufacturer_id = ?", manufacturer.ID).
			Update("manufacturer_id", 0).
			Error

		if err != nil {
			return err
		}

		return tx.Delete(&manufacturer).Error
	})
}

// UserRepository

func (t *userRepository) GetByName(name string) (storage.User, error) {
	var user storage.User

	err := t.db.Preload(clause.Associations).
		Where("name = ?", name).
		First(&user).
		Error

	return user, err
}

func (t *userRepository) Save(user storage.User) error {
	if string(user.Password) == "" {
		return errors.New("password can not be empty")
	}

	if _, err := bcrypt.Cost(user.Password); err != nil {
		pw, err := bcrypt.GenerateFromPassword(user.Password, bcrypt.DefaultCost)

		if err != nil {
			return err
		}

		user.Password = pw
	}

	return t.db.Clauses(
		clause.OnConflict{
			UpdateAll: true,
		}).
		Save(&user).
		Error
}

// ReportRepository

func (t *reportRepository) Paginate(options storage.QueryOptions) ([]storage.Report, error) {
	var reports []storage.Report

	err := t.paginate(options).
		Find(&reports).
		Error

	return reports, err
}

// Save will overwrite the report for its owning Asset
// All CVE associations will be updated
func (t *reportRepository) Save(report storage.Report) error {
	cve := report.Cves

	return t.db.Transaction(func(tx *gorm.DB) error {
		err := t.db.FirstOrCreate(&report, storage.Report{AssetID: report.AssetID}).
			Error

		if err != nil {
			return err
		}

		return t.db.Model(&report).
			Clauses(
				clause.OnConflict{
					DoNothing: true,
					Columns: []clause.Column{
						{
							Name: "CVE",
						},
					},
				}).
			Association("Cves").
			Append(cve)
	})

}
