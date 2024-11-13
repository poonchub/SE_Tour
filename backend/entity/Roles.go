package entity

import "gorm.io/gorm"

type Roles struct {
	gorm.Model
	RoleName	string

	Employees	[]Employees `gorm:"foreignKey:RoleID"`
}