package entity

import (
	"time"

	"gorm.io/gorm"
)

type Promotions struct {
	gorm.Model
	PromotionName 	string
	DiscountPercentage	float32
	ValidFrom	time.Time
	ValidUntil	time.Time
	Minimum_Price	float32

	PromotionStatusID	uint

	Bookings	[]Bookings	`gorm:"foreignKey:PromotionID"`
}