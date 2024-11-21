package entity

import "gorm.io/gorm"

type TravelInsurances struct {
	gorm.Model
	Price			float32
	CoverageDetail	string

	ProviderID	uint
	Provider	Providers	`gorm:"foreignKey:ProviderID"`

	InsuranceParticipants	[]InsuranceParticipants	`gorm:"foreignKey:TravelInsuranceID"`
}