package sqlite

import (
	"github.com/cmp307/assetman/pkg/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (t *repository) PaginateByName(needle string, options storage.QueryOptions) ([]storage.Asset, error) {
	var assets []storage.Asset

	tx := t.db.Preload(clause.Associations).
		Order(options.Order)

	if len(needle) > 0 {
		tx.Where("name like ?", "%"+needle+"%")
	}

	if options.Limit > 0 {
		tx.Limit(options.Limit)
	}

	if options.Offset > 0 {
		tx.Offset(options.Offset)
	}

	err := tx.Find(&assets).
		Error

	return assets, err
}

func (t *repository) Add(asset interface{}) error {
	switch v := asset.(type) {
	case storage.HardwareAsset:
		return t.db.Create(&v).Error
	case storage.SoftwareAsset:
		return t.db.Create(&v).Error
	default:
		panic("invalid type")
	}

	return nil
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

func (t *repository) GetByName(name string) (storage.User, error) {
	var user storage.User

	err := t.db.Where("name = ?", name).
		First(&user).
		Error

	return user, err
}
