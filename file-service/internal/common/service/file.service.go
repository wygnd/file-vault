package service

import (
	"github.com/wygnd/file-vault/file-service/internal/common/dto"
	"github.com/wygnd/file-vault/file-service/internal/common/helpers"
	"github.com/wygnd/file-vault/file-service/internal/common/mappers"
	"github.com/wygnd/file-vault/file-service/internal/common/models"
	filerepository "github.com/wygnd/file-vault/file-service/internal/common/repository"
	"github.com/wygnd/file-vault/file-service/pkg/minio"
	helpers2 "github.com/wygnd/file-vault/file-service/pkg/minio/helpers"
)

type FileService interface {
	Upload(file dto.UploadFile) (*dto.FileResponseDTO, error)
	GetByID(id string) (*dto.FileDetailResponseDTO, error)
	Delete(id string) error
	ListByFolderID(folderId string) ([]*dto.FileResponseDTO, error)
}

type fileService struct {
	repo  filerepository.FileRepository
	minio minio.Client
}

func NewFileService(repo filerepository.FileRepository, minio minio.Client) FileService {
	return &fileService{
		repo:  repo,
		minio: minio,
	}
}

// Upload загружает файл в сервис
func (service *fileService) Upload(file dto.UploadFile) (*dto.FileResponseDTO, error) {

	// Создаем объект в S3
	objectId, err := service.minio.CreateOne(helpers2.FileDataType{
		FileName: file.Name,
		MimeType: file.MimeType,
		Size:     file.Size,
		Data:     file.Data,
	})

	if err != nil {
		return nil, err
	}

	// Формируем объект для сохранения в БД
	record := &models.File{
		Name:       file.Name,
		StorageKey: objectId,
		MimeType:   file.MimeType,
		Size:       file.Size,
		HashSha256: helpers.GenerateHash(file.Data),
		FolderID:   &file.FolderId,
		OwnerID:    file.OwnerId,
	}

	// Создаем запись в БД
	err = service.repo.Create(record)

	if err != nil {
		_ = service.minio.DeleteOne(objectId)

		return nil, err
	}

	return mappers.ToFileResponseDTO(record), nil
}

// GetByID получает ссылку на файл по ID
func (service *fileService) GetByID(id string) (*dto.FileDetailResponseDTO, error) {

	record, errGetRecord := service.repo.GetById(id)

	if errGetRecord != nil {
		return nil, errGetRecord
	}

	url, errGetUrl := service.minio.GetOne(record.StorageKey)

	if errGetUrl != nil {
		return nil, errGetUrl
	}

	return &dto.FileDetailResponseDTO{
		FileResponseDTO: *mappers.ToFileResponseDTO(record),
		URL:             url,
	}, nil
}

// Delete удаляет объект из БД и S3
func (service *fileService) Delete(id string) error {

	// Получаем запись из БД
	record, err := service.repo.GetById(id)

	// Если не удалось получить: выдаем ошибку
	if err != nil {
		return err
	}

	// Удаляем запись из БД
	err = service.repo.Delete(id)

	// Не получилось: выдаем ошибку
	if err != nil {
		return err
	}

	// Удаляем из хранилища
	err = service.minio.DeleteOne(record.StorageKey)

	// Не получилось: выдаем ошибку
	if err != nil {
		return err
	}

	return nil
}

// ListByFolderID получает список файлов в конкретной папке
func (service *fileService) ListByFolderID(folderId string) ([]*dto.FileResponseDTO, error) {
	records, err := service.repo.ListByFolderID(folderId)

	if err != nil {
		return nil, err
	}

	dtoList := make([]*dto.FileResponseDTO, 0, len(records))
	for _, record := range records {
		dtoList = append(dtoList, mappers.ToFileResponseDTO(&record))
	}

	return dtoList, nil
}
