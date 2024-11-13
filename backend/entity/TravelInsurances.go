package entity

import "gorm.io/gorm"

type TravelInsurances struct {
	gorm.Model
	Price			float32
	CoverageDetail	string

	BookingID	uint
	Booking		Bookings	`gorm:"foreignKey:BookingID"`

	ProviderID	uint
	Provider	Providers	`gorm:"foreignKey:ProviderID"`
}