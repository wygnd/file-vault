package db

import (
	"github.com/wygnd/file-vault/file-service/internal/common/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	return db.AutoMigrate(&models.File{}, &models.Folder{})
}
