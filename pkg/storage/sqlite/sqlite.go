package sqlite

import (
	"github.com/cmp307/assetman/pkg/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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

	err := t.db.
		Preload("Manufacturer").
		Preload("SoftwareAsset").
		Preload("HardwareAsset").
		Find(&assets).Error

	return assets, err
}

func (t *repository) GetById(id uint) storage.Asset {
	var asset storage.Asset

	t.db.Preload("Manufacturer").
		Preload("SoftwareAsset").
		Preload("HardwareAsset").
		Where("id = ?", id).
		First(&asset)

	return asset
}

func (t *repository) GetAllByNameContains(needle string) ([]storage.Asset, error) {
	var assets []storage.Asset

	err := t.db.
		Preload("Manufacturer").
		Preload("SoftwareAsset").
		Preload("HardwareAsset").
		Where("name like '%?%'", needle).
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

func (t *repository) GetByName(name string) (storage.User, error) {
	var user storage.User
	err := t.db.Where("name = ?", name).First(&user).Error

	return user, err
}
