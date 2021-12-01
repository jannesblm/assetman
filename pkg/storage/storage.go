package storage

import (
	"gorm.io/gorm"
	"reflect"
	"time"
)

const TypeHardware = "hardware"
const TypeSoftware = "software"

type BasicAsset interface {
	GetID() uint
}

// Asset holds general information about an asset and is either linked to a specific software or
// hardware asset.
type Asset struct {
	gorm.Model
	Name           string `gorm:"index"`
	Description    string
	ManufacturerID uint
	Manufacturer   Manufacturer `gorm:"foreignKey:ID;references:ManufacturerID"`
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
	MAC               string
	IP                string
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

func (a Asset) GetID() uint {
	return a.ID
}

func (a Asset) Type() reflect.Type {
	return reflect.TypeOf(a.Polymorphic())
}

func (a Asset) Polymorphic() BasicAsset {
	switch a.AssetType {
	case TypeHardware:
		return a.HardwareAsset
	case TypeSoftware:
		return a.SoftwareAsset
	default:
		panic("AssetType invalid")
	}
}

func (a *HardwareAsset) GetID() uint {
	return a.ID
}

func (a *SoftwareAsset) GetID() uint {
	return a.ID
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
	CountAll() int64
	CountHardware() int64
	CountSoftware() int64
	// GetAll retrieves all stored assets from the database.
	GetAll() ([]Asset, error)
	GetAllManufacturers() ([]Manufacturer, error)
	GetById(uint) Asset
	PaginateByName(string, QueryOptions) ([]Asset, error)
	PaginateByTypeAndName(string, string, QueryOptions) ([]Asset, error)
	Save(Asset) error
}

type UserRepository interface {
	GetByName(name string) (User, error)
}
