package entity

import (
	"time"

	"gorm.io/gorm"
)

type Accommodations struct {
	gorm.Model
	CheckInDate		time.Time
	CheckOutDate	time.Time

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`

	HotelID		uint
	Hotel		Hotels	`gorm:"foreignKey:HotelID"`

	BookingRooms	[]BookingRooms	`gorm:"foreignKey:AccommodationID"`
}