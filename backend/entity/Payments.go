package entity

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	Amount	float32

	PaymentMethodID	uint
	PaymentMethod	PaymentMethods	`gorm:"foreignKey:PaymentMethodID"`

	BookingID		uint
	Booking			Bookings	`gorm:"foreignKey:BookingID"`

	PaymentStatusID	uint
	PaymentStatus	PaymentStatuses	`gorm:"foreignKey:PaymentStatusID"`

	SalesReportID	uint
	SalesReport		SalesReports	`gorm:"foreignKey:SalesReportID"`
}