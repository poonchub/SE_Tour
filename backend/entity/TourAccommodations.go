package entity

import "gorm.io/gorm"

type TourAccommodations struct {
	gorm.Model
	AccommodationID	uint
	Accommodation	Accommodations	`gorm:"foreignKey:AccommodationID"`

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`
}