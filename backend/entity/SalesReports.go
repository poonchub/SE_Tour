package entity

import "gorm.io/gorm"

type SalesReports struct {
	gorm.Model
	ReportName	string
	Data		string

	Payments	[]Payments	`gorm:"foreignKey:SalesReportID"`
}