package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type MinioConfig struct {
	Endpoint   string
	AccessKey  string
	SecretKey  string
	BucketName string
	UseSSL     bool
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

type Config struct {
	Port     string
	Minio    MinioConfig
	Database DatabaseConfig
}

var AppConfig *Config

func LoadConfig() {
	_ = godotenv.Load()

	// Собираем объект с env переменными
	AppConfig = &Config{
		Port: getEnv("FILE_SERVICE_PORT", "5000"),
		Minio: MinioConfig{
			Endpoint:   getEnv("MINIO_ENDPOINT", "localhost:9000"),
			AccessKey:  getEnv("MINIO_ACCESS_KEY", ""),
			SecretKey:  getEnv("MINIO_SECRET_KEY", ""),
			BucketName: getEnv("MINIO_BUCKET_NAME", "defaultBucket"),
			UseSSL:     getEnvAsBool("MINIO_USE_SSL", false),
		},
		Database: DatabaseConfig{
			Host:     getEnv("FILE_SERVICE_DATABASE_HOST", "localhost"),
			Port:     getEnv("FILE_SERVICE_DATABASE_PORT", "5432"),
			User:     getEnv("FILE_SERVICE_DATABASE_USER", "postgres"),
			Password: getEnv("FILE_SERVICE_DATABASE_PASSWORD", "postgres"),
			Name:     getEnv("FILE_SERVICE_DATABASE_NAME", "file_service"),
		},
	}
}

// Получение env переменной
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

// Получение численной evn переменной
func getEnvAsInt(key string, defaultValue int) int {
	valueString := getEnv(key, "")

	if value, err := strconv.Atoi(valueString); err == nil {
		return value
	}

	return defaultValue
}

// Получение булевой env переменной
func getEnvAsBool(key string, defaultValue bool) bool {
	valueString := getEnv(key, "")

	if value, err := strconv.ParseBool(valueString); err == nil {
		return value
	}

	return defaultValue
}

// Получение списка из env переменной
func getEnvAsSlice(key string, separator string, defaultValue []string) []string {
	valueString := getEnv(key, "")

	if valueString == "" {
		return defaultValue
	}

	return strings.Split(valueString, separator)
}
