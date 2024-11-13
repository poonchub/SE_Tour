package entity

import "gorm.io/gorm"

type PaymentMethods struct {
	gorm.Model
	MethodName	string

	Payments	[]Payments	`gorm:"foreignKey:PaymentMethodID"`
}