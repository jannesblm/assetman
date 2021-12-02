package sqlite

import (
	"errors"
	"github.com/cmp307/assetman/pkg/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
)

type Repository interface {
	Connect() (*gorm.DB, error)
}

type repository struct {
	db *gorm.DB
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

func Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}

func NewAssetRepository(db *gorm.DB) storage.AssetRepository {
	db.AutoMigrate(
		storage.Asset{},
		storage.Manufacturer{},
		storage.HardwareAsset{},
		storage.SoftwareAsset{},
		storage.User{},
	)

	return &assetRepository{
		repository{db},
	}
}

func NewManufacturerRepository(db *gorm.DB) storage.ManufacturerRepository {
	db.AutoMigrate(
		storage.Asset{},
		storage.Manufacturer{},
		storage.HardwareAsset{},
		storage.SoftwareAsset{},
		storage.User{},
	)

	return &manufRepository{
		repository{db},
	}
}

func (t repository) paginate(options storage.QueryOptions) *gorm.DB {
	tx := t.db.Preload(clause.Associations).
		Order(options.Order)

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
		Where("id = ?", id).
		First(&asset)

	return asset
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

	err := tx.Find(&assets).
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
	err := t.db.
		Model(reflect.New(asset.Type()).Interface()).
		Clauses(
			clause.OnConflict{
				UpdateAll: true,
			}).
		Create(asset.Polymorphic()).
		Error

	if err != nil {
		return err
	}

	asset.AssetID = asset.Polymorphic().GetID()

	return t.db.
		Session(&gorm.Session{FullSaveAssociations: true}).
		Omit("HardwareAsset").
		Omit("SoftwareAsset").
		Save(&asset).
		Error
}

func (t *assetRepository) Delete(asset storage.Asset) error {
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
}

// ManufacturerRepository

func (t *manufRepository) GetById(id uint) storage.Manufacturer {
	var manufacturer storage.Manufacturer

	t.db.Preload(clause.Associations).
		Where("id = ?", id).
		First(&manufacturer)

	return manufacturer
}

func (t *manufRepository) Paginate(options storage.QueryOptions) ([]storage.Manufacturer, error) {
	var manufacturers []storage.Manufacturer

	tx := t.paginate(options)

	if len(options.Query) > 0 {
		tx.Where(options.QueryField+" like ?", "%"+options.Query+"%")
	}

	err := tx.Find(&manufacturers).
		Error

	return manufacturers, err
}

func (t *manufRepository) CountAll() int64 {
	var count int64
	t.db.Model(&storage.Manufacturer{}).Count(&count)

	return count
}

// UserRepository

func (t *repository) GetByName(name string) (storage.User, error) {
	var user storage.User

	err := t.db.Where("name = ?", name).
		First(&user).
		Error

	return user, err
}
