package entity

import "gorm.io/gorm"

type BookingStatus struct {
	gorm.Model
	StatusName	string
}