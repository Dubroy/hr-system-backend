package models

import (
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"size:100;not null"`
	Code      string         `gorm:"size:50;uniqueIndex;not null"`
	Users     []User         `gorm:"foreignKey:DepartmentID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
