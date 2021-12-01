package sqlite

import (
	"errors"
	"github.com/cmp307/assetman/pkg/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"reflect"
)

type repository struct {
	db *gorm.DB
}

func Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
}

func NewRepository(db *gorm.DB) *repository {
	/*db.AutoMigrate(
		storage.Asset{},
		storage.HardwareAsset{},
		storage.SoftwareAsset{},
		storage.User{},
	)*/

	return &repository{
		db,
	}
}

func (t *repository) GetAll() ([]storage.Asset, error) {
	var assets []storage.Asset

	err := t.db.Preload(clause.Associations).
		Find(&assets).
		Error

	return assets, err
}

func (t *repository) GetById(id uint) storage.Asset {
	var asset storage.Asset

	t.db.Preload(clause.Associations).
		Where("id = ?", id).
		First(&asset)

	return asset
}

func (t *repository) paginate(options storage.QueryOptions) *gorm.DB {
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

func (t *repository) PaginateByName(needle string, options storage.QueryOptions) ([]storage.Asset, error) {
	var assets []storage.Asset

	tx := t.paginate(options)

	if len(needle) > 0 {
		tx.Where("name like ?", "%"+needle+"%")
	}

	err := tx.Find(&assets).
		Error

	return assets, err
}

func (t *repository) PaginateByTypeAndName(typ string, needle string, options storage.QueryOptions) ([]storage.Asset, error) {
	var assets []storage.Asset

	tx := t.paginate(options)

	if typ != storage.TypeHardware && typ != storage.TypeSoftware {
		return []storage.Asset{}, errors.New("invalid type")
	}

	tx.Where("asset_type = ?", typ)

	if len(needle) > 0 {
		tx.Where("name like ?", "%"+needle+"%")
	}

	err := tx.Find(&assets).
		Error

	return assets, err
}

func (t *repository) GetAllManufacturers() ([]storage.Manufacturer, error) {
	var manufacturers []storage.Manufacturer
	err := t.db.Find(&manufacturers).Error

	return manufacturers, err
}

func (t *repository) CountAll() int64 {
	var count int64
	t.db.Model(&storage.Asset{}).Count(&count)

	return count
}

func (t *repository) CountSoftware() int64 {
	var count int64

	t.db.Model(&storage.Asset{}).
		Where("asset_type = ?", storage.TypeSoftware).
		Count(&count)

	return count
}

func (t *repository) CountHardware() int64 {
	var count int64

	t.db.Model(&storage.Asset{}).
		Where("asset_type = ?", storage.TypeHardware).
		Count(&count)

	return count
}

func (t *repository) Save(asset storage.Asset) error {
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

func (t *repository) GetByName(name string) (storage.User, error) {
	var user storage.User

	err := t.db.Where("name = ?", name).
		First(&user).
		Error

	return user, err
}
