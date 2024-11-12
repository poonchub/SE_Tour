package entity

import "gorm.io/gorm"

type RoomTypes struct {
	gorm.Model
	TypeName	string

	Accommodations	[]Accommodations `gorm:"foreignKey:RoomTypeID"`
}