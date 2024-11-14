package entity

import "gorm.io/gorm"

type BookingRooms struct {
	gorm.Model
	RoomQuantity	int

	AccommodationID	uint
	Accommodation	Accommodations	`gorm:"foreignKey:AccommodationID"`

	RoomTypeID	uint
	RoomType	RoomTypes	`gorm:"foreignKey:RoomTypeID"`
}