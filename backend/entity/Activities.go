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

	LocationID   	uint
    Location    	Locations   `gorm:"foreignKey:LocationID"`

	TourActivities 	[]TourActivities	`gorm:"foreignKey:ActivityID"`
}