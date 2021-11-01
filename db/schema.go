package db

import (
	"gorm.io/gorm"
	"time"
)

type States struct {
	gorm.Model
	UserID       uint `gorm:"not null"`
	CourseID     uint `gorm:"not null"`
	Complete       bool `gorm:"type:bool"`
	Progress     int
	ActualPoints int
	Game string
	StartTime    time.Time `gorm:"type:TIME"`
	CompleteTime time.Time `gorm:"type:TIME"`
}

type Results struct {
	gorm.Model
	StatesID uint `gorm:"not null"`
	Key string
	Points   int
	Value  string
	States   States `gorm:"foreignKey:StatesID"`
}
