package entity

import (
	"gorm.io/gorm"
)

type Activities struct {
	gorm.Model
	ActivityName	string
	Description		string

	LocationID   	uint
    Location    	Locations   `gorm:"foreignKey:LocationID"`

	ScheduleActivities	[]ScheduleActivities	`gorm:"foreignKey:ActivityID"`
}