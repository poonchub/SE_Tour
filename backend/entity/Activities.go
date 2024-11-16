package entity

import (
	"time"

	"gorm.io/gorm"
)

type Activities struct {
	gorm.Model
	ActivityName	string
	Description		string
	DateTime		time.Time

	LocationID   	uint
    Location    	Locations   `gorm:"foreignKey:LocationID"`

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`
}