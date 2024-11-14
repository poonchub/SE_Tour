package entity

import (
	"gorm.io/gorm"
)

type EmployeeSchedules struct {
	gorm.Model

	TourScheduleID	uint
	TourSchedule	TourSchedules	`gorm:"foreignKey:TourScheduleID"`

	EmployeeID		uint
	Employee		Employees	`gorm:"foreignKey:EmployeeID"`
}