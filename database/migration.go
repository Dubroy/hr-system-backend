package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Migration struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);uniqueIndex"` // 修改這裡
	AppliedAt time.Time
}

func RunMigrations(db *gorm.DB) error {
	// 創建 migrations 表來記錄已執行的 migration
	err := db.AutoMigrate(&Migration{})
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %v", err)
	}

	// 讀取 migrations 目錄中的所有 SQL 文件
	files, err := filepath.Glob("database/migrations/*.sql")
	if err != nil {
		return fmt.Errorf("failed to read migration files: %v", err)
	}

	// 排序文件名
	sort.Strings(files)

	// 執行每個 migration 文件
	for _, file := range files {
		filename := filepath.Base(file)
		
		// 檢查 migration 是否已執行
		var count int64
		db.Model(&Migration{}).Where("name = ?", filename).Count(&count)
		if count > 0 {
			log.Printf("Migration %s already applied, skipping...", filename)
			continue
		}

		// 讀取 SQL 文件內容
		content, err := os.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", filename, err)
		}

		// 分割多個 SQL 語句（假設使用分號分隔）
		statements := strings.Split(string(content), ";")

		// 開始事務
		tx := db.Begin()

		// 執行每個 SQL 語句
		for _, stmt := range statements {
			stmt = strings.TrimSpace(stmt)
			if stmt == "" {
				continue
			}

			if err := tx.Exec(stmt).Error; err != nil {
				tx.Rollback()
				return fmt.Errorf("failed to execute migration %s: %v", filename, err)
			}
		}

		// 記錄已執行的 migration
		if err := tx.Create(&Migration{
			Name:      filename,
			AppliedAt: time.Now(),
		}).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to record migration %s: %v", filename, err)
		}

		// 提交事務
		if err := tx.Commit().Error; err != nil {
			return fmt.Errorf("failed to commit migration %s: %v", filename, err)
		}

		log.Printf("Successfully applied migration: %s", filename)
	}

	return nil
} 