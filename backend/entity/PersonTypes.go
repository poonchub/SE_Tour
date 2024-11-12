package entity

import "gorm.io/gorm"

type PersonTypes struct {
	gorm.Model
	TypeName	string

	BookingDetails 	[]BookingDetails `gorm:"foreignKey:PersonTypeID"`
	TourPrices		[]TourPrices `gorm:"foreignKey:PersonTypeID"`
}