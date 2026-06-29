package repository

import (
	"github.com/wygnd/file-vault/file-service/internal/common/models"
	"gorm.io/gorm"
)

type FileRepository interface {
	Create(file *models.File) error
	GetById(ID string) (*models.File, error)
	Delete(ID string) error
	ListByFolderID(folderID string) ([]models.File, error)
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) FileRepository {
	return &fileRepository{db: db}
}

// Create создает запись в БД
func (repo *fileRepository) Create(file *models.File) error {
	result := repo.db.Create(file)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetById получает запись по ID
func (repo *fileRepository) GetById(ID string) (*models.File, error) {
	var file models.File

	result := repo.db.First(&file, ID)

	if result.Error != nil {
		return nil, result.Error
	}

	return &file, nil
}

// Delete удаляет запись по ID
func (repo *fileRepository) Delete(ID string) error {
	var file models.File

	result := repo.db.Delete(&file, ID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// ListByFolderID получает список файлов по ID папки
func (repo *fileRepository) ListByFolderID(folderID string) ([]models.File, error) {
	var files []models.File

	result := repo.db.Where("folder_id = ?", folderID).Find(&files)

	if result.Error != nil {
		return nil, result.Error
	}

	return files, nil
}
