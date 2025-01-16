package api_test

import (
	"bytes"
	"encoding/json"
	v1 "hr-system-backend/api/v1"
	"hr-system-backend/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// 修改 mock 定義
type mockEmployeeService struct {
	createFunc func(employee *models.Employee) error
}

// 實現 services.EmployeeService 介面
func (m *mockEmployeeService) CreateEmployee(employee *models.Employee) error {
	if m.createFunc != nil {
		return m.createFunc(employee)
	}
	return nil
}

func (m *mockEmployeeService) GetEmployee(id uint) (*models.Employee, error) {
	return nil, nil
}

func (m *mockEmployeeService) ListEmployees(page, pageSize int) ([]models.Employee, int64, error) {
	return nil, 0, nil
}

func TestCreateEmployee(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name       string
		reqBody    map[string]interface{}
		setupMock  func(*mockEmployeeService)
		wantStatus int
		wantBody   map[string]interface{}
	}{
		{
			name: "成功創建員工",
			reqBody: map[string]interface{}{
				"name":       "張小明",
				"email":      "test@example.com",
				"department": "IT",
				"position":   "工程師",
			},
			setupMock: func(m *mockEmployeeService) {
				m.createFunc = func(employee *models.Employee) error {
					return nil
				}
			},
			wantStatus: http.StatusCreated,
			wantBody: map[string]interface{}{
				"message": "員工創建成功",
			},
		},
		{
			name: "缺少必要欄位",
			reqBody: map[string]interface{}{
				"name": "張小明",
			},
			setupMock: func(m *mockEmployeeService) {
				m.createFunc = func(employee *models.Employee) error {
					return nil
				}
			},
			wantStatus: http.StatusBadRequest,
			wantBody: map[string]interface{}{
				"error": "缺少必要欄位",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 設置 mock service
			mockService := &mockEmployeeService{}
			tt.setupMock(mockService)

			// 創建 handler
			handler := v1.NewEmployeeHandler(mockService)

			// 創建 router
			router := gin.New()
			router.POST("/api/v1/employees", handler.Create)

			// 創建請求體
			jsonData, _ := json.Marshal(tt.reqBody)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bytes.NewBuffer(jsonData))
			req.Header.Set("Content-Type", "application/json")

			// 記錄響應
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// 驗證狀態碼
			assert.Equal(t, tt.wantStatus, w.Code)

			// 驗證響應內容
			var response map[string]interface{}
			err := json.Unmarshal(w.Body.Bytes(), &response)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantBody, response)
		})
	}
} 