package entity

import "gorm.io/gorm"

type Providers struct {
	gorm.Model
	ProviderName	string

	TravelInsurances	[]TravelInsurances	`gorm:"foreignKey:ProviderID"`
}