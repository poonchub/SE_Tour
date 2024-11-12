package entity

import (
	"time"

	"gorm.io/gorm"
)

type StaffSchedules struct {
	gorm.Model
	ShiftStart	time.Time
	ShiftEnd	time.Time

	StaffID		uint
	TourPackageID	uint
}