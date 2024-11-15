package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transportations struct {
	gorm.Model
	DepartureTime	time.Time
	ArrivalTime		time.Time

	VehicleID	uint
	Vehicle		Vehicles	`gorm:"foreignKey:VehicleID"`

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`

	LocationID		uint
	Location		Locations		`gorm:"foreignKey:LocationID"`
}