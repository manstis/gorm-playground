package main

import (
	"gorm.io/gorm"
)

type People []Person
type Person struct {
	gorm.Model
	AddressID *int
	Address   Address
}

type Address struct {
	gorm.Model
	AddressType  string
	AddressLines []AddressLine `gorm:"foreignKey:AddressID"`
}

type AddressLine struct {
	AddressID *int   `gorm:"primaryKey;index"`
	Text      string `gorm:"primaryKey;not null"`
}
