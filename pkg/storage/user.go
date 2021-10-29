package storage

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint
	Name     string
	Password []byte `gorm:"size:60"`
}

type UserRepository interface {
	GetByName(name string) (User, error)
}
