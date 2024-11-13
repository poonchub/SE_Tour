package entity

import "gorm.io/gorm"

type TourActivities struct {
	gorm.Model
	ActivityID	uint
	Activity	Activities	`gorm:"foreignKey:ActivityID"`

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`
}