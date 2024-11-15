package entity

import "gorm.io/gorm"

type Locations struct {
	gorm.Model
	LocationName	string

	Activities	[]Activities	`gorm:"foreignKey:LocationID"`

	Transportations	[]Transportations	`gorm:"foreignKey:LocationID"`
}