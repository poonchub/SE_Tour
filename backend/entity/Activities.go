package entity

import (
	"time"

	"gorm.io/gorm"
)

type Activities struct {
	gorm.Model
	ActivityName	string
	Description		string
	StartTime		time.Time
	EndTime			time.Time
	Location		string

	TourActivities 	[]TourActivities	`gorm:"foreignKey:ActivityID"`
}