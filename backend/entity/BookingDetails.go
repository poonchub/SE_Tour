package entity

import "gorm.io/gorm"

type BookingDetails struct {
	gorm.Model
	Quantity	int

	BookingID	uint
	Booking		Bookings	`gorm:"foreignKey:BookingID"`

	TourPriceID	uint
	TourPrice	TourPrices	`gorm:"foreignKey:TourPriceID"`
}