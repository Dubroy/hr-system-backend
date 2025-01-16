package repositories

import (
	"hr-system-backend/models"

	"gorm.io/gorm"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

func (r *EmployeeRepository) Create(employee *models.Employee) error {
	return r.db.Create(employee).Error
}

func (r *EmployeeRepository) FindByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	err := r.db.Preload("Department").Preload("Position").First(&employee, id).Error
	return &employee, err
}

func (r *EmployeeRepository) List(page, pageSize int) ([]models.Employee, int64, error) {
	var employees []models.Employee
	var total int64

	query := r.db.Model(&models.Employee{})
	query.Count(&total)

	err := query.Offset((page - 1) * pageSize).
		Limit(pageSize).
		Preload("Department").
		Preload("Position").
		Find(&employees).Error

	return employees, total, err
} 