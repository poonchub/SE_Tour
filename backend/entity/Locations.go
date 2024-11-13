package entity

import "gorm.io/gorm"

type Locations struct {
	gorm.Model
	LocationName	string

	ProvinceID	uint
	Province	Provinces	`gorm:"foreignKey:ProvinceID"`

	Activities	[]Activities	`gorm:"foreignKey:LocationID"`

	Transportations	[]Transportations	`gorm:"foreignKey:LocationID"`
}