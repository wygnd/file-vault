package repository

import (
	"github.com/wygnd/file-vault/file-service/internal/common/dto"
	"github.com/wygnd/file-vault/file-service/internal/common/models"
	"gorm.io/gorm"
)

type FolderRepository interface {
	Create(folder *models.Folder) (*models.Folder, error)
	Update(ID string, dto dto.UpdateFolderDTO) (*models.Folder, error)
	Delete(ID string) error
	GetChildren(ownerID string, parentID string) ([]models.Folder, error)
}

type folderRepository struct {
	db *gorm.DB
}

func NewFolderRepository(db *gorm.DB) FolderRepository {
	return &folderRepository{db: db}
}

// Create создает запись в БД
func (repo *folderRepository) Create(folder *models.Folder) (*models.Folder, error) {
	result := repo.db.Create(folder)

	if result.Error != nil {
		return nil, result.Error
	}

	return folder, nil
}

// Update обновляет запись в БД
func (repo *folderRepository) Update(ID string, dto dto.UpdateFolderDTO) (*models.Folder, error) {
	var folder models.Folder

	// Мапим только то, что предали в DTO. Безопасное обновление
	result := repo.db.Model(&folder).Where("id = ?", ID).Updates(map[string]interface{}{
		"name": dto.Name,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return &folder, nil
}

// Delete удаляет запись из БД
func (repo *folderRepository) Delete(ID string) error {
	var folder models.Folder

	result := repo.db.Delete(&folder, ID)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

// GetChildren получает дочерние папки
func (repo *folderRepository) GetChildren(ownerID string, parentID string) ([]models.Folder, error) {
	var folders []models.Folder

	result := repo.db.Where("parentId = ?", ownerID).Where("id = ?", parentID).Find(&folders)

	if result.Error != nil {
		return nil, result.Error
	}

	return folders, nil
}
