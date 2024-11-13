package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transportations struct {
	gorm.Model
	DepartureTime	time.Time

	VehicleTypeID	uint
	VehicleType		VehicleTypes	`gorm:"foreignKey:VehicleTypeID"`

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`

	LocationID		uint
	Location		Locations		`gorm:"foreignKey:LocationID"`
}