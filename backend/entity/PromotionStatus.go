package entity

import "gorm.io/gorm"

type PromotionStatus struct {
	gorm.Model
	StatusName	string

	Promotions	[]Promotions `gorm:"foreignKey:PromotionStatusID"`
}