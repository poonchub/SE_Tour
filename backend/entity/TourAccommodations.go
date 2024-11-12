package entity

import "gorm.io/gorm"

type TourAccommodations struct {
	gorm.Model
	AccommodationID	uint
	TourPackageID	uint
}