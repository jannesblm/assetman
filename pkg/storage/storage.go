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
	TypeName       string
	ManufacturerID uint
	Manufacturer   Manufacturer `gorm:"foreignKey:ID;references:ManufacturerID"`
	PurchaseDate   time.Time
	Note           string
	AssetID        uint           `gorm:"index"`
	AssetType      string         `gorm:"index"`
	Report         Report         `gorm:"foreignKey:AssetID"`
	HardwareAsset  *HardwareAsset `gorm:"foreignKey:ID;references:AssetID"`
	SoftwareAsset  *SoftwareAsset `gorm:"foreignKey:ID;references:AssetID"`
}

// HardwareAsset holds fields specific to hardware assets
type HardwareAsset struct {
	ID                uint   `gorm:"primaryKey"`
	Asset             Asset  `gorm:"polymorphic:Asset;polymorphicValue:hardware"`
	ModelName         string `gorm:"index"`
	InternalID        string
	MAC               string
	IP                string
	Location          string
	WarrantyInfo      string
	InstalledSoftware []Asset `gorm:"many2many:hardware_software"`
}

// SoftwareAsset holds fields specific to software assets
type SoftwareAsset struct {
	ID          uint  `gorm:"primaryKey"`
	Asset       Asset `gorm:"polymorphic:Asset;polymorphicValue:software"`
	Version     string
	LicenseType string
	LicenseKey  string
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

// GetID implementation to conform with BasicAsset
func (a *HardwareAsset) GetID() uint {
	return a.ID
}

func (a *SoftwareAsset) GetID() uint {
	return a.ID
}

// Manufacturer has been normalised to its own entity
type Manufacturer struct {
	gorm.Model
	Name   string  `gorm:"uniqueIndex"`
	Assets []Asset `gorm:"foreignKey:ID;references:ID"`
}

type QueryOptions struct {
	Query      string
	QueryField string
	Limit      int
	Offset     int
	Order      string
}

type User struct {
	gorm.Model
	Name     string
	Group    string
	Password []byte `json:"-" gorm:"size:60"`
}

type Cve struct {
	ReportID uint   `gorm:"primaryKey"`
	CVE      string `gorm:"index;unique"`
}

type Report struct {
	ID         uint `gorm:"primaryKey"`
	AssetID    uint `gorm:"unique;foreignKey:ID;references:AssetID"`
	ReportedAt time.Time
	Cves       []Cve `gorm:"many2many:report_cve"`
}

type AssetRepository interface {
	CountAll() int64
	CountHardware() int64
	CountSoftware() int64
	GetById(uint) Asset
	Paginate(string, QueryOptions) ([]Asset, error)
	GetAllSoftware() ([]Asset, error)
	Save(Asset) (uint, error)
	Delete(Asset) error
}

type ManufacturerRepository interface {
	GetAll() ([]Manufacturer, error)
	CountAll() int64
	Paginate(QueryOptions) ([]Manufacturer, error)
	Delete(Manufacturer) error
}

type UserRepository interface {
	GetByName(name string) (User, error)
	Save(User) error
}

type ReportRepository interface {
	Paginate(QueryOptions) ([]Report, error)
	Save(Report) error
}
