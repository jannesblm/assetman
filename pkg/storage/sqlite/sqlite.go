package sqlite

import (
	"assetman/pkg/storage"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository() *repository {
	db, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})

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
