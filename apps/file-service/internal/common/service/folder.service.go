package service

import (
	"github.com/wygnd/file-vault/file-service/internal/common/dto"
	"github.com/wygnd/file-vault/file-service/internal/common/mappers"
	"github.com/wygnd/file-vault/file-service/internal/common/models"
	"github.com/wygnd/file-vault/file-service/internal/common/repository"
)

type FolderService interface {
	Create(folder dto.CreateFolderDTO) (*dto.FolderResponseDTO, error)
	Update(ID string, fields *dto.UpdateFolderDTO) (*dto.FolderResponseDTO, error)
	Delete(ID string) error
	GetChildren(ownerID string, parentID *string) (*dto.FolderChildrenResponseDTO, error)
}

type folderService struct {
	folderRepo repository.FolderRepository
	fileRepo   repository.FileRepository
}

func NewFolderService(folderRepo repository.FolderRepository, fileRepo repository.FileRepository) FolderService {
	return &folderService{folderRepo: folderRepo, fileRepo: fileRepo}
}

// Create создает запись в БД
func (service *folderService) Create(folder dto.CreateFolderDTO) (*dto.FolderResponseDTO, error) {
	record, err := service.folderRepo.Create(&models.Folder{
		Name:     folder.Name,
		ParentID: folder.ParentID,
		OwnerID:  folder.OwnerID,
	})

	if err != nil {
		return nil, err
	}

	return mappers.ToFolderResponseDTO(record), nil
}

// Delete удаляет запись из БД
func (service *folderService) Delete(ID string) error {
	err := service.folderRepo.Delete(ID)

	if err != nil {
		return err
	}

	return nil
}

// Update обновляет данные
func (service *folderService) Update(ID string, fields *dto.UpdateFolderDTO) (*dto.FolderResponseDTO, error) {

	result, err := service.folderRepo.Update(ID, *fields)

	if err != nil {
		return nil, err
	}

	return &dto.FolderResponseDTO{
		ID:        result.ID,
		Name:      result.Name,
		ParentID:  result.ParentID,
		OwnerID:   result.OwnerID,
		CreatedAt: result.CreatedAt,
	}, nil

}

// GetChildren получает список дочерних папок и файлов
func (service *folderService) GetChildren(ownerID string, parentID *string) (*dto.FolderChildrenResponseDTO, error) {
	folders, err := service.folderRepo.GetChildren(ownerID, *parentID)

	if err != nil {
		return nil, err
	}

	files, err := service.fileRepo.ListByFolderID(*parentID)

	if err != nil {
		return nil, err
	}

	dtoFolderList := make([]*dto.FolderResponseDTO, 0, len(folders))

	for _, record := range folders {
		dtoFolderList = append(dtoFolderList, mappers.ToFolderResponseDTO(&record))
	}

	dtoFileList := make([]*dto.FileResponseDTO, 0, len(files))

	for _, file := range files {
		dtoFileList = append(dtoFileList, mappers.ToFileResponseDTO(&file))
	}

	return &dto.FolderChildrenResponseDTO{
		Folders: dtoFolderList,
		Files:   dtoFileList,
	}, nil
}
