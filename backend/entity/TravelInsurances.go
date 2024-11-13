package entity

import "gorm.io/gorm"

type TravelInsurances struct {
	gorm.Model
	InsuranceProvider string
	CoverageDetail    string
}