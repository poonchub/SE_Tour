package entity

import "gorm.io/gorm"

type Roles struct {
	gorm.Model
	RoleName	string

	Staffs		[]Staffs `gorm:"foreignKey:RoleID"`
}