package entity

import "gorm.io/gorm"

type Meals struct {
	gorm.Model
	MenusDetail	string

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`

	MealTypeID		uint
	MealType		MealTypes		`gorm:"foreignKey:MealTypeID"`
}