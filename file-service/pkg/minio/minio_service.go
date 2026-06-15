package minio

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/wygnd/file-vault/file-service/internal/common/config"
	"github.com/wygnd/file-vault/file-service/pkg/minio/helpers"
)

// CreateOne Создает объект в бакете Minio
func (m *minioClient) CreateOne(file helpers.FileDataType) (string, error) {
	objectID := uuid.New().String()

	reader := bytes.NewReader(file.Data)

	_, err := m.mc.PutObject(context.Background(), config.AppConfig.Minio.BucketName, objectID, reader, int64(len(file.Data)), minio.PutObjectOptions{})

	if err != nil {
		return "", fmt.Errorf("Ошибка при создании объекта %s: %v", file.FileName, err)
	}

	_, err = m.mc.PresignedGetObject(context.Background(), config.AppConfig.Minio.BucketName, objectID, time.Second*24*60*60, nil)

	if err != nil {
		return "", fmt.Errorf("Ошибка при генерации ссылки %s: %v", file.FileName, err)
	}

	return objectID, nil
}

// CreateMany создает несколько объектов в хранилище Minio
func (m *minioClient) CreateMany(data map[string]helpers.FileDataType) ([]string, error) {
	urlList := make([]string, 0, len(data))

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	// Создание канала для передачи URL-адресов с размером, равным кол-ву передаваемых данных
	urlChannel := make(chan string, len(data))

	// WaitGroup, чтобы ожидать завершения всех горутин
	var wg sync.WaitGroup

	for objectID, file := range data {
		// Увеличение счетчика WaitGroup перед запуском горутины
		wg.Add(1)

		go func(objectID string, file helpers.FileDataType) {
			// Уменьшение счетчика WaitGroup после завершения горутины
			defer wg.Done()

			// Пытаемся создать объект в бакете
			_, errCreateObject := m.mc.PutObject(ctx, config.AppConfig.Minio.BucketName, objectID, bytes.NewReader(file.Data), int64(len(file.Data)), minio.PutObjectOptions{})

			// Если произошла ошибка, отменяем операцию
			if errCreateObject != nil {
				cancel()
				return
			}

			// Получаем публичный URL файла
			url, errGetUrl := m.mc.PresignedGetObject(ctx, config.AppConfig.Minio.BucketName, objectID, time.Second*24*60*60, nil)

			// Не удалось получить URL: отменяем операцию
			if errGetUrl != nil {
				cancel()
				return
			}

			urlChannel <- url.String()
		}(objectID, file)
	}

	// Ожидания завершения всех горутин и закрытие канала со списком URL
	go func() {
		wg.Wait()         // Блокируем, пока счетчик WaitingGroup не станет равным 0
		close(urlChannel) // Закрываем канал с URL адресами после завершения всех горутин
	}()

	// Собираем все URL адреса

	for url := range urlChannel {
		urlList = append(urlList, url)
	}

	return urlList, nil
}

// GetOne получает один объект по ID
func (m *minioClient) GetOne(objectID string) (string, error) {
	url, err := m.mc.PresignedGetObject(context.Background(), config.AppConfig.Minio.BucketName, objectID, time.Second*24*60*60, nil)

	if err != nil {
		return "", fmt.Errorf("ошибка получения URL: %v", err)
	}

	return url.String(), nil
}

// GetMany получает несколько объектов по IDs
func (m *minioClient) GetMany(objectIDs []string) ([]string, error) {
	urlChannel := make(chan string, len(objectIDs))
	errChannel := make(chan helpers.OperationError, len(objectIDs))

	var wg sync.WaitGroup

	_, cancel := context.WithCancel(context.Background())

	defer cancel()

	for _, objectID := range objectIDs {
		wg.Add(1)

		go func(objectID string) {
			defer wg.Done()

			url, err := m.GetOne(objectID)

			if err != nil {
				errChannel <- helpers.OperationError{
					ObjectId: objectID,
					Error:    err,
				}
				cancel()
				return
			}

			urlChannel <- url
		}(objectID)
	}

	go func() {
		wg.Wait()
		close(urlChannel)
		close(errChannel)
	}()

	var urlList []string
	var errList []error

	for url := range urlChannel {
		urlList = append(urlList, url)
	}

	for err := range errChannel {
		errList = append(errList, err.Error)
	}

	if len(errList) > 0 {
		return nil, fmt.Errorf("ошибки при получении файлов: %v", errList)
	}

	return urlList, nil
}

// DeleteOne удаляем объект по ID
func (m *minioClient) DeleteOne(objectID string) error {
	err := m.mc.RemoveObject(context.Background(), config.AppConfig.Minio.BucketName, objectID, minio.RemoveObjectOptions{})

	if err != nil {
		return err
	}

	return nil
}

// DeleteMany удаляем несколько объектов по IDs
func (m *minioClient) DeleteMany(objectIDs []string) error {
	errChannel := make(chan helpers.OperationError, len(objectIDs))

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	for _, objectID := range objectIDs {
		wg.Add(1)

		go func(objectID string) {
			defer wg.Done()

			err := m.mc.RemoveObject(ctx, config.AppConfig.Minio.BucketName, objectID, minio.RemoveObjectOptions{})

			if err != nil {
				errChannel <- helpers.OperationError{
					ObjectId: objectID,
					Error:    err,
				}
				cancel()
			}

		}(objectID)
	}

	go func() {
		wg.Wait()
		close(errChannel)
	}()

	var errList []error

	for err := range errChannel {
		errList = append(errList, err.Error)
	}

	if len(errList) > 0 {
		return fmt.Errorf("ошибка при удалении объектов: %v", errList)
	}

	return nil
}
