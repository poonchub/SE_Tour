package entity

import "gorm.io/gorm"

type Provinces struct {
	gorm.Model
	ProvinceName	string

	Locations	[]Locations	`gorm:"foreignKey:ProvinceID"`
}