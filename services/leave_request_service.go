package services

import (
	"errors"
	"hr-system-backend/models"
	"hr-system-backend/repositories"
)

type LeaveRequestService struct {
	repo *repositories.LeaveRequestRepository
}

func NewLeaveRequestService(repo *repositories.LeaveRequestRepository) *LeaveRequestService {
	return &LeaveRequestService{repo: repo}
}

func (s *LeaveRequestService) CreateLeaveRequest(leave *models.LeaveRequest) error {
	// 計算請假天數
	days := int(leave.EndDate.Sub(leave.StartDate).Hours() / 24) + 1
	if days < 1 {
		return errors.New("end date must be after start date")
	}
	leave.TotalDays = days

	return s.repo.Create(leave)
}

func (s *LeaveRequestService) GetLeaveRequest(id uint) (*models.LeaveRequest, error) {
	return s.repo.FindByID(id)
}

func (s *LeaveRequestService) ListLeaveRequests(employeeID uint, status string, page, pageSize int) ([]models.LeaveRequest, int64, error) {
	return s.repo.List(employeeID, status, page, pageSize)
}

func (s *LeaveRequestService) ApproveLeaveRequest(id uint, approverID uint) error {
	return s.repo.UpdateStatus(id, "approved", approverID)
}

func (s *LeaveRequestService) RejectLeaveRequest(id uint, approverID uint) error {
	return s.repo.UpdateStatus(id, "rejected", approverID)
} 