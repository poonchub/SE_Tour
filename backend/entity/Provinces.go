package entity

import "gorm.io/gorm"

type Provinces struct {
	gorm.Model
	ProvinceName	string

	TourPackages	[]TourPackages	`gorm:"foreignKey:ProvinceID"`
	Locations		[]Locations		`gorm:"foreignKey:ProvinceID"`
}