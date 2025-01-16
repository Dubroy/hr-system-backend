package models

import (
	"time"

	"gorm.io/gorm"
)

type Position struct {
	ID        uint           `gorm:"primaryKey"`
	Title     string         `gorm:"size:100;not null"`
	Level     int           `gorm:"not null"`
	Users     []User         `gorm:"foreignKey:PositionID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
