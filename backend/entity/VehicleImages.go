package entity

import "gorm.io/gorm"

type VehicleImages struct {
	gorm.Model
	FilePath string
	VehicleID uint
	Vehicle Vehicles `gorm:"foreignKey:VehicleID"`
}