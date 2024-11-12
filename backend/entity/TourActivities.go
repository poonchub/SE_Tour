package entity

import "gorm.io/gorm"

type TourActivities struct {
	gorm.Model
	ActivityID	uint
	TourPackageID	uint
}