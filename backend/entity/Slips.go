package entity

import "gorm.io/gorm"

type Slips struct {
	gorm.Model
	FilePath	string
	PaymentID	uint
}