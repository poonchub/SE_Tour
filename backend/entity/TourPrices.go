package entity

import "gorm.io/gorm"

type TourPrices struct {
	gorm.Model
	Price 	float32

	TourPackageID	uint
	PersonTypeID	uint
}