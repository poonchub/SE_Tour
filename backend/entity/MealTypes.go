package entity

import "gorm.io/gorm"

type MealTypes struct {
	gorm.Model
	TypeName	string

	Meals	[]Meals	`gorm:"foreignKey:MealTypeID"`
}