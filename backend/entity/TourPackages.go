package entity

import (
	"gorm.io/gorm"
)

type TourPackages struct {
	gorm.Model
	TourName	string
	PackageCode	string
	Description	string
	Duration	string

	TourPrices	[]TourPrices	`gorm:"foreignKey:TourPackageID"`

	TourSchedules	[]TourSchedules	`gorm:"foreignKey:TourPackageID"`

	TourImages	[]TourImages	`gorm:"foreignKey:TourPackageID"`

	Transportations	[]Transportations	`gorm:"foreignKey:TourPackageID"`

	Meals		[]Meals			`gorm:"foreignKey:TourPackageID"`
}