package entity

import "gorm.io/gorm"

type Payments struct {
	gorm.Model
	Amount	float32

	PaymentMethodID	uint
	BookingID		uint
	PaymentStatusID	uint
}