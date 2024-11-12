package entity

import "gorm.io/gorm"

type SalesReports struct {
	gorm.Model
	Data	string
}