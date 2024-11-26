package entity

import (
	"time"

	"gorm.io/gorm"
)

type ScheduleActivities struct {
	gorm.Model
	DateTime	time.Time

	ActivityID	uint
	Activity	Activities	`gorm:"foreignKey:ActivityID"`

	TourScheduleID 	uint
	TourSchedule	TourSchedules	`gorm:"foreignKey:TourScheduleID"`
}