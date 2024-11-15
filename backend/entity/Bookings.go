package entity

import "gorm.io/gorm"

type Bookings struct {
	gorm.Model
	TotalPrice		float32

	CustomerID		uint
	Customer		Customers	`gorm:"foreignKey:CustomerID"`

	TourScheduleID 	uint
	TourSchedule	TourSchedules	`gorm:"foreignKey:TourScheduleID"`

	BookingStatusID	uint
	BookingStatus	BookingStatuses	`gorm:"foreignKey:BookingStatusID"`

	PromotionID		uint
	Promotion		Promotions	`gorm:"foreignKey:PromotionID"`

	BookingDetails	[]BookingDetails	`gorm:"foreignKey:BookingID"`

	Payments		*Payments			`gorm:"foreignKey:BookingID"`
}