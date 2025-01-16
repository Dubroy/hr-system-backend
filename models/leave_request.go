package models

import (
	"time"

	"gorm.io/gorm"
)

type LeaveRequestInput struct {
	EmployeeID uint   `json:"employee_id"`
	LeaveType  string `json:"leave_type"`
	StartDate  int64  `json:"start_date"` // Unix timestamp
	EndDate    int64  `json:"end_date"`   // Unix timestamp
	Reason     string `json:"reason"`
}

type LeaveRequest struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	EmployeeID uint          `json:"employee_id" gorm:"not null"`
	Employee   Employee      `json:"employee" gorm:"foreignKey:EmployeeID"`
	LeaveType  string        `json:"leave_type" gorm:"size:50;not null"`
	StartDate  time.Time     `json:"start_date" gorm:"not null"`
	EndDate    time.Time     `json:"end_date" gorm:"not null"`
	TotalDays  int           `json:"total_days" gorm:"not null"`
	Reason     string        `json:"reason" gorm:"type:text"`
	Status     string        `json:"status" gorm:"size:20;default:pending"`
	ApprovedBy *uint         `json:"approved_by"`
	Approver   *Employee     `json:"approver" gorm:"foreignKey:ApprovedBy"`
	ApprovedAt *time.Time    `json:"approved_at"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
} 