package sqlite

import (
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

	return &repository{
		db,
	}
}
