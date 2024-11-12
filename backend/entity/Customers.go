package entity

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	UserName 	string
	FirstName 	string
	LastName	string
	Email		string
	Password	string
	PhoneNumber	string
}