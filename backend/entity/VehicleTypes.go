package entity

import "gorm.io/gorm"

type VehicleTypes struct {
	gorm.Model
	TypeName	string

	Transportations	[]Transportations `gorm:"foreignKey:VehicleTypeID"`
}