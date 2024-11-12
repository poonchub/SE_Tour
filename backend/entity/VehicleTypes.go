package entity

import "gorm.io/gorm"

type VehicleTypes struct {
	gorm.Model
	TypeName	string

	Transpportations	[]Transpportations `gorm:"foreignKey:VehicleTypeID"`
}