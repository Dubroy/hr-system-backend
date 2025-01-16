package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint           `gorm:"primaryKey"`
	Name         string         `gorm:"size:100;not null"`
	Email        string         `gorm:"size:100;uniqueIndex;not null"`
	DepartmentID uint
	Department   Department
	PositionID   uint
	Position     Position
	Status       string         `gorm:"size:20;default:'active'"` // active, inactive, terminated
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
