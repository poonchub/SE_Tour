package entity

import "gorm.io/gorm"

type Customers struct {
	gorm.Model
	UserName 	string
	FirstName 	string
	LastName	string
	Email		string
	Password	string
	PhoneNumber	string
	ProfilePath	string

	Bookings	[]Bookings	`gorm:"foreignKey:CustomerID"`
}