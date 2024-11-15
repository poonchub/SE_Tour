package entity

import "gorm.io/gorm"

type TourDescriptions struct {
	gorm.Model
	Intro	string
	PackageDetail	string
	TripHighlight	string
	PlacesHighlight	string

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`
}