package main

import (
	"fmt"
	"hr-system-backend/config"
	"hr-system-backend/database"
	"log"

	v1 "hr-system-backend/api/v1"
	"hr-system-backend/repositories"
	"hr-system-backend/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRoutes(r *gin.Engine, db *gorm.DB) {
	// 初始化 repositories
	employeeRepo := repositories.NewEmployeeRepository(db)
	leaveRepo := repositories.NewLeaveRequestRepository(db)

	// 初始化 services
	employeeService := services.NewEmployeeService(employeeRepo)
	leaveService := services.NewLeaveRequestService(leaveRepo)

	// 初始化 handlers
	employeeHandler := v1.NewEmployeeHandler(employeeService)
	leaveHandler := v1.NewLeaveRequestHandler(leaveService)

	// API 路由
	apiV1 := r.Group("/api/v1")
	{
		// 員工相關路由
		employees := apiV1.Group("/employees")
		{
			employees.POST("/", employeeHandler.Create)
			employees.GET("/", employeeHandler.List)
			employees.GET("/:id", employeeHandler.Get)
		}

		// 請假相關路由
		leaves := apiV1.Group("/leaves")
		{
			leaves.POST("/", leaveHandler.Create)
			leaves.GET("/", leaveHandler.List)
			leaves.GET("/:id", leaveHandler.Get)
			leaves.PUT("/:id/approve", leaveHandler.Approve)
			leaves.PUT("/:id/reject", leaveHandler.Reject)
		}
	}
}

func main() {
	// 加載配置
	cfg := config.LoadConfig()
	fmt.Printf("Loaded Config: %+v\n", cfg)

	// 連接數據庫
	db := database.ConnectDB(cfg)

	// 執行 migrations
	if err := database.RunMigrations(db); err != nil {
		log.Fatal("Failed to run migrations:", err)
	}

	// 初始化 Gin 引擎
	r := gin.Default()

	// 健康檢查路由
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 設定 API 路由
	setupRoutes(r, db)

	// 啟動伺服器
	log.Printf("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
