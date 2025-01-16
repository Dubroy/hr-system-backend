package services

import (
	"hr-system-backend/models"
	"hr-system-backend/repositories"
)

type EmployeeService struct {
	repo *repositories.EmployeeRepository
}

func NewEmployeeService(repo *repositories.EmployeeRepository) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (s *EmployeeService) CreateEmployee(employee *models.Employee) error {
	// 可以在這裡添加業務邏輯驗證
	return s.repo.Create(employee)
}

func (s *EmployeeService) GetEmployee(id uint) (*models.Employee, error) {
	return s.repo.FindByID(id)
}

func (s *EmployeeService) ListEmployees(page, pageSize int) ([]models.Employee, int64, error) {
	return s.repo.List(page, pageSize)
} 