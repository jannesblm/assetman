package storage

import (
	"gorm.io/gorm"
	"time"
)

// Asset holds general information about an asset and is either linked to a specific software or
// hardware asset.
type Asset struct {
	gorm.Model
	Name           string `gorm:"index"`
	Description    string
	ManufacturerID uint
	Manufacturer   Manufacturer
	PurchaseDate   time.Time
	Note           string
	AssetID        uint           `gorm:"index"`
	AssetType      string         `gorm:"index"`
	HardwareAsset  *HardwareAsset `gorm:"foreignKey:ID;references:AssetID"`
	SoftwareAsset  *SoftwareAsset `gorm:"foreignKey:ID;references:AssetID"`
}

type HardwareAsset struct {
	ID                uint   `gorm:"primaryKey"`
	Asset             Asset  `gorm:"polymorphic:Asset;polymorphicValue:hardware"`
	ModelName         string `gorm:"index"`
	InternalID        string
	MAC               uint64
	IP                uint64
	Location          string
	WarrantyInfo      string
	InstalledSoftware []SoftwareAsset `gorm:"many2many:hardware_software"`
}

type SoftwareAsset struct {
	ID          uint  `gorm:"primaryKey"`
	Asset       Asset `gorm:"polymorphic:Asset;polymorphicValue:software"`
	Version     string
	LicenseType string
}

type Manufacturer struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

type User struct {
	gorm.Model
	Name     string
	Password []byte `gorm:"size:60"`
}

type QueryOptions struct {
	Limit  int
	Offset int
	Order  string `default:"id desc"`
}

type AssetRepository interface {
	// GetAll retrieves all stored assets from the database.
	GetAll() ([]Asset, error)
	CountAll() int64
	GetById(uint) Asset
	GetAllManufacturers() ([]Manufacturer, error)
	PaginateByName(string, QueryOptions) ([]Asset, error)
}

type UserRepository interface {
	GetByName(name string) (User, error)
}
