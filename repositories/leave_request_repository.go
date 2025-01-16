package repositories

import (
	"hr-system-backend/models"
	"time"

	"gorm.io/gorm"
)

type LeaveRequestRepository struct {
	db *gorm.DB
}

func NewLeaveRequestRepository(db *gorm.DB) *LeaveRequestRepository {
	return &LeaveRequestRepository{db: db}
}

func (r *LeaveRequestRepository) Create(leave *models.LeaveRequest) error {
	return r.db.Create(leave).Error
}

func (r *LeaveRequestRepository) FindByID(id uint) (*models.LeaveRequest, error) {
	var leave models.LeaveRequest
	err := r.db.Preload("Employee").Preload("Approver").First(&leave, id).Error
	return &leave, err
}

func (r *LeaveRequestRepository) List(employeeID uint, status string, page, pageSize int) ([]models.LeaveRequest, int64, error) {
	var leaves []models.LeaveRequest
	var total int64

	query := r.db.Model(&models.LeaveRequest{})

	if employeeID > 0 {
		query = query.Where("employee_id = ?", employeeID)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	err := query.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Preload("Employee").
		Preload("Approver").
		Find(&leaves).Error

	return leaves, total, err
}

func (r *LeaveRequestRepository) UpdateStatus(id uint, status string, approverID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		leave := models.LeaveRequest{}
		if err := tx.First(&leave, id).Error; err != nil {
			return err
		}

		now := time.Now()
		updates := map[string]interface{}{
			"status":      status,
			"approved_by": approverID,
			"approved_at": now,
		}

		return tx.Model(&leave).Updates(updates).Error
	})
} 