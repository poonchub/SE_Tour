package entity

import "gorm.io/gorm"

type PaymentMethods struct {
	gorm.Model
	MethodName	string

	
}