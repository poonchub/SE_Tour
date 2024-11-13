package entity

import (
	"gorm.io/gorm"
)

type EmployeeSchedules struct {
	gorm.Model
	EmployeeID		uint
	Employee		Employees	`gorm:"foreignKey:EmployeeID"`

	TourPackageID	uint
	TourPackage		TourPackages	`gorm:"foreignKey:TourPackageID"`
}