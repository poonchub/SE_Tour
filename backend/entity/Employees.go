package entity

import "gorm.io/gorm"

type Employees struct {
	gorm.Model
	UserName 	string
	FirstName	string
	LastName	string
	Email		string
	Password	string
	PhoneNumber	string
	ProfilePath	string

	RoleID		uint
	Role		Roles	`gorm:"foreignKey:RoleID"`

	EmployeeSchedules	[]EmployeeSchedules	`gorm:"foreignKey:EmployeeID"`
}