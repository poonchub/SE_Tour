package entity

import "gorm.io/gorm"

type Bookings struct {
	gorm.Model
	TotalPrice		float32

	CustomerID		uint
	TourPackageID 	uint
	BookingStatusID	uint
	PromotionID		uint
}