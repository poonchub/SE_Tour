package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transpportations struct {
	gorm.Model
	DepartureTime	time.Time

	VehicleTypeID	uint
	VehicleType		VehicleTypes	`gorm:"foreignKey:VehicleTypeID"`

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`
}