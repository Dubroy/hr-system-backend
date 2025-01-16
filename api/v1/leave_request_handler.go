package v1

import (
	"hr-system-backend/models"
	"hr-system-backend/services"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type LeaveRequestHandler struct {
	service *services.LeaveRequestService
}

func NewLeaveRequestHandler(service *services.LeaveRequestService) *LeaveRequestHandler {
	return &LeaveRequestHandler{service: service}
}

func (h *LeaveRequestHandler) Create(c *gin.Context) {
	var input models.LeaveRequestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	leave := &models.LeaveRequest{
		EmployeeID: input.EmployeeID,
		LeaveType:  input.LeaveType,
		StartDate:  time.Unix(input.StartDate, 0),
		EndDate:    time.Unix(input.EndDate, 0),
		Reason:     input.Reason,
	}

	if err := h.service.CreateLeaveRequest(leave); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, leave)
}

func (h *LeaveRequestHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	leave, err := h.service.GetLeaveRequest(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Leave request not found"})
		return
	}

	c.JSON(http.StatusOK, leave)
}

func (h *LeaveRequestHandler) List(c *gin.Context) {
	employeeID, _ := strconv.ParseUint(c.DefaultQuery("employee_id", "0"), 10, 32)
	status := c.DefaultQuery("status", "")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	leaves, total, err := h.service.ListLeaveRequests(uint(employeeID), status, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": leaves,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
}

func (h *LeaveRequestHandler) Approve(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// TODO: 從 JWT 獲取 approverID
	approverID := uint(1) // 暫時硬編碼

	if err := h.service.ApproveLeaveRequest(uint(id), approverID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leave request approved"})
}

func (h *LeaveRequestHandler) Reject(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// TODO: 從 JWT 獲取 approverID
	approverID := uint(1) // 暫時硬編碼

	if err := h.service.RejectLeaveRequest(uint(id), approverID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Leave request rejected"})
} 