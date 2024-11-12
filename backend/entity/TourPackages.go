package entity

import (
	"time"

	"gorm.io/gorm"
)

type TourPackages struct {
	gorm.Model
	TourName	string
	Description	string
	StartDate	time.Time
	EndDate		time.Time
	MaxParticipants	int
}