package storage

import "gorm.io/gorm"

type Asset struct {
	gorm.Model
	ID             uint
	Name           string
	TypeID         uint
	Type           Type
	ManufacturerID uint
	Manufacturer   Manufacturer
}

type Type struct {
	gorm.Model
	ID   uint
	Name string
}

type Manufacturer struct {
	gorm.Model
	ID   uint
	Name string
}

type AssetRepository interface {
	// GetAll retrieves all stored assets from the database.
	GetAll() ([]Asset, error)
}
