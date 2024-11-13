package entity

import "gorm.io/gorm"

type Providers struct {
	gorm.Model
	ProviderName	string
	LogoPath		string

	TravelInsurances	[]TravelInsurances	`gorm:"foreignKey:ProviderID"`
}