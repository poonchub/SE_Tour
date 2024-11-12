package entity

import (
	"time"

	"gorm.io/gorm"
)

type Accommodations struct {
	gorm.Model
	HotelName	string
	CheckInDate	time.Time
	CheckOutDate	time.Time

	RoomTypeID	uint
}