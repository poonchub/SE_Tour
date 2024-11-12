package entity

import "gorm.io/gorm"

type BookingDetails struct {
	gorm.Model
	Quantity	int

	BookingID	uint
	PersonTypeID	uint
}