package models

import "gorm.io/gorm"

type File struct {
	gorm.Model

	ID         string  `gorm:"type:uuid;primaryKey;default:gorm_random_uuid()"`
	Name       string  `gorm:"type:varchar(100);not null"`
	StorageKey string  `gorm:"column:storage_key;type:varchar(150);not null"`
	MimeType   string  `gorm:"column:mime_type;type:varchar(50);not null"`
	Size       int64   `gorm:"type:bigint;not null"`
	HashSha256 string  `gorm:"column:hash_sha256;type:varchar(255);not null"`
	FolderID   *string `gorm:"column:folder_id;type:varchar(255)"`
	OwnerID    string  `gorm:"column:owner_id;type:varchar(255);not null"`
}
