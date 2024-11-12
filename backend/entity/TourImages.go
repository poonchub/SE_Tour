package entity

import "gorm.io/gorm"

type TourImages struct {
	gorm.Model
	FilePath	string

	TourPackageID	uint
}