package sqlite

import (
	"errors"
	"github.com/cmp307/assetman/pkg/storage"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
)

type Repository interface {
	Connect() (*gorm.DB, error)
}

type DB struct {
	*gorm.DB
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

func Connect(path string) (*DB, error) {
	g, err := gorm.Open(sqlite.Open(path), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

	if err != nil {
		return &DB{}, err
	}

	return &DB{g}, nil
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

func NewAssetRepository(db *DB) storage.AssetRepository {
	return &assetRepository{
		repository{db},
	}
}

func NewManufacturerRepository(db *DB) storage.ManufacturerRepository {
	return &manufRepository{
		repository{db},
	}
}

func NewUserRepository(db *DB) storage.UserRepository {
	return &userRepository{
		repository{db},
	}
}

func NewReportRepository(db *DB) storage.ReportRepository {
	return &reportRepository{
		repository{db},
	}
}

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

		tx.Where(options.QueryField+" like ?", "%"+options.Query+"%")
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

func (t *assetRepository) Save(asset storage.Asset) error {
	return t.db.Transaction(func(tx *gorm.DB) error {
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
			Session(&gorm.Session{FullSaveAssociations: false}).
			Save(&asset).
			Error
	})
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
		tx.Where(options.QueryField+" like ?", "%"+options.Query+"%")
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
