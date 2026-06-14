package minio

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/wygnd/file-vault/file-service/internal/common/config"
	"github.com/wygnd/file-vault/file-service/pkg/minio/helpers"
)

type Client interface {
	Init() error
	CreateOne(file helpers.FileDataType) (string, error)
	CreateMany(data map[string]helpers.FileDataType) ([]string, error)
	GetOne(objectID string) (string, error)
	GetMany(objectIDs []string) ([]string, error)
	DeleteOne(objectID string) error
	DeleteMany(objectIDs []string) error
}

type minioClient struct {
	mc *minio.Client
}

func NewMinioClient() Client {
	return &minioClient{}
}

func (m *minioClient) Init() error {
	ctx := context.Background()

	client, err := minio.New(config.AppConfig.Minio.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AppConfig.Minio.AccessKey, config.AppConfig.Minio.SecretKey, ""),
		Secure: config.AppConfig.Minio.UseSSL,
	})

	if err != nil {
		return err
	}

	// Подключаемся к Minio
	m.mc = client

	// Проверяем, существует ли бакет
	bucketExists, err := m.mc.BucketExists(ctx, config.AppConfig.Minio.BucketName)

	if err != nil {
		return err
	}

	// Если не существует, пытаемся создать бакет
	if !bucketExists {
		err := m.mc.MakeBucket(ctx, config.AppConfig.Minio.BucketName, minio.MakeBucketOptions{})

		if err != nil {
			return nil
		}
	}

	return nil
}
