package entity

import "gorm.io/gorm"

type BookingDetails struct {
	gorm.Model
	Quantity	int

	BookingID	uint
	Booking		Bookings	`gorm:"foreignKey:BookingID"`

	PersonTypeID	uint
	PersonType	PersonTypes	`gorm:"foreignKey:PersonTypeID"`
}