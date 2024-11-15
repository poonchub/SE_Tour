package entity

import "gorm.io/gorm"

type Vehicles struct {
	gorm.Model
	VehicleName	string

	VehicleTypeID	uint
	VehicleType		VehicleTypes	`gorm:"foreignKey:VehicleTypeID"`

	Transportations	[]Transportations	`gorm:"foreignKey:VehicleID"`
}