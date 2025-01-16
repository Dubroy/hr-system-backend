package models

import (
	"time"

	"gorm.io/gorm"
)

type Employee struct {
	ID           uint           `gorm:"primaryKey"`
	EmployeeCode string         `gorm:"size:50;uniqueIndex"`
	Name         string         `gorm:"size:100;not null"`
	Email        string         `gorm:"size:100;uniqueIndex"`
	PhoneNumber  string         `gorm:"size:20"`
	DepartmentID uint
	Department   Department
	PositionID   uint
	Position     Position
	Salary       float64        `gorm:"type:decimal(10,2);not null"`
	Status       string         `gorm:"size:20;default:'active'"`
	JoinDate     time.Time
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
} 