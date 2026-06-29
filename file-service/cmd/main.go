package main

import (
	"log"
	"net"

	gen "github.com/wygnd/file-vault/file-service/gen/file"
	"github.com/wygnd/file-vault/file-service/internal/common/config"
	"github.com/wygnd/file-vault/file-service/internal/common/db"
	"github.com/wygnd/file-vault/file-service/internal/common/repository"
	"github.com/wygnd/file-vault/file-service/internal/common/service"
	grpcFileService "github.com/wygnd/file-vault/file-service/pkg/grpc"
	"github.com/wygnd/file-vault/file-service/pkg/minio"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {
	// Загрузка .env файла
	config.LoadConfig()

	// Подключаемся к Minio
	minioClient := minio.NewMinioClient()
	err := minioClient.Init()

	if err != nil {
		log.Fatalf("Ошибка инициализации Minio: %v", err)
	}

	// Подключаемся к БД
	database, err := db.Init()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	// Подключаем Репозитории
	fileRepository := repository.NewFileRepository(database)
	folderRepository := repository.NewFolderRepository(database)

	// Подключаем Сервисы
	fileService := service.NewFileService(fileRepository, minioClient)
	folderService := service.NewFolderService(folderRepository, fileRepository)

	// gRPC server
	grpcServer := grpc.NewServer()
	fileGrpcService := grpcFileService.NewFileGrpcService(fileService, folderService)
	gen.RegisterFileServiceServer(grpcServer, fileGrpcService)

	// healthcheck server
	healthServer := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthServer)
	healthServer.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)

	listener, err := net.Listen("tcp", ":"+config.AppConfig.Port)

	if err != nil {
		log.Fatalf("Ошибка запуска listener: %v", err)
	}

	log.Printf("gRPC сервер запущен на порту %s", config.AppConfig.Port)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка запуска gRPC сервера: %v", err)
	}
}
