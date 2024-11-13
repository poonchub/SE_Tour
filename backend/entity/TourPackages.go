package entity

import (
	"time"

	"gorm.io/gorm"
)

type TourPackages struct {
	gorm.Model
	TourName	string
	Description	string
	StartDate	time.Time
	EndDate		time.Time
	MaxParticipants	int

	TourPrices	[]TourPrices	`gorm:"foreignKey:TourPackageID"`
	Bookings	[]Bookings	`gorm:"foreignKey:TourPackageID"`
	TourImages	[]TourImages	`gorm:"foreignKey:TourPackageID"`
	StaffSchedules	[]StaffSchedules	`gorm:"foreignKey:TourPackageID"`
	TourAccommodations	[]TourAccommodations	`gorm:"foreignKey:TourPackageID"`
	TourActivities	[]TourActivities	`gorm:"foreignKey:TourPackageID"`
	Transpportations	[]Transpportations	`gorm:"foreignKey:TourPackageID"`
}