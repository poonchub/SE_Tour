package entity

import "gorm.io/gorm"

type Admins struct {
	gorm.Model
	UserName 	string
	FirstName	string
	LastName	string
	Email		string
	Password	string
}