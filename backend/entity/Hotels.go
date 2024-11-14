package entity

import "gorm.io/gorm"

type Hotels struct {
	gorm.Model
	HotelName string

	Accommodations []Accommodations `gorm:"foreignKey:HotelID"`
}