package entity

import "gorm.io/gorm"

type Staffs struct {
	gorm.Model
	UserName 	string
	FirstName	string
	LastName	string
	Email		string
	Password	string
	PhoneNumber	string

	RoleID		uint
}