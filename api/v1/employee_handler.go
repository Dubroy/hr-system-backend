package v1

import (
	"hr-system-backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 首先定義一個介面
type EmployeeServiceInterface interface {
	CreateEmployee(employee *models.Employee) error
	GetEmployee(id uint) (*models.Employee, error)
	ListEmployees(page, pageSize int) ([]models.Employee, int64, error)
}

type EmployeeHandler struct {
	service EmployeeServiceInterface  // 使用介面而不是具體類型
}

func NewEmployeeHandler(service EmployeeServiceInterface) *EmployeeHandler {
	return &EmployeeHandler{service: service}
}

func (h *EmployeeHandler) Create(c *gin.Context) {
	var employee models.Employee
	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateEmployee(&employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, employee)
}

func (h *EmployeeHandler) Get(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	employee, err := h.service.GetEmployee(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	c.JSON(http.StatusOK, employee)
}

func (h *EmployeeHandler) List(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))

	employees, total, err := h.service.ListEmployees(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": employees,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"page_size": pageSize,
		},
	})
} 