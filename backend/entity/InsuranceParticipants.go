package entity

import "gorm.io/gorm"

type InsuranceParticipants struct {
	gorm.Model
	IdCardNumber	string
	FirstName		string
	LastName		string
	Age				int
	PhoneNumber		string

	BookingID		uint
	Booking			Bookings	`gorm:"foreignKey:BookingID"`

	TravelInsuranceID	uint
	TravelInsurance		TravelInsurances `gorm:"foreignKey:TravelInsuranceID"`
}