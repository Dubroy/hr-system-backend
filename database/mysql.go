package database

import (
	"fmt"
	"hr-system-backend/config"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	var db *gorm.DB
	var err error
	
	// 重試次數
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to database")
			return db
		}
		
		retrySeconds := time.Duration(i+1) * 2 
		log.Printf("Failed to connect to database. Retrying in %v... (Attempt %d/%d)\n", retrySeconds, i+1, maxRetries)
		time.Sleep(retrySeconds)
	}

	log.Fatal("Failed to connect to database after multiple attempts:", err)
	return nil
} 