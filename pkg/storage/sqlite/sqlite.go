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
	db.AutoMigrate(
		&storage.Asset{},
		storage.User{},
	)

	return &repository{
		db,
	}
}

func (t *repository) GetAll() ([]storage.Asset, error) {
	var assets []storage.Asset
	err := t.db.Find(&assets).Error

	return assets, err
}

func (t *repository) GetByName(name string) (storage.User, error) {
	var user storage.User
	err := t.db.Where("name = ?", name).First(&user).Error

	return user, err
}
