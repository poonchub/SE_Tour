package entity

import (
	"time"

	"gorm.io/gorm"
)

type Transpportations struct {
	gorm.Model
	DepartureTime	time.Time

	VehicleTypeID	uint
	TourPackageID	uint
}