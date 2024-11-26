package entity

import (
	"time"

	"gorm.io/gorm"
)

type TourSchedules struct {
	gorm.Model
	StartDate	time.Time
	EndDate		time.Time
	AvailableSlots	int

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`

	TourScheduleStatusID	uint
	TourScheduleStatus	TourScheduleStatuses	`gorm:"foreignKey:TourScheduleStatusID"`

	EmployeeSchedules	[]EmployeeSchedules	`gorm:"foreignKey:TourScheduleID"`

	Bookings	[]Bookings	`gorm:"foreignKey:TourScheduleID"`

	ScheduleActivities	[]ScheduleActivities	`gorm:"foreignKey:TourScheduleID"`
}