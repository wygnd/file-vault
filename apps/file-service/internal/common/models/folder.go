package models

import (
	"time"

	"gorm.io/gorm"
)

type Folder struct {
	ID        string  `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Name      string  `gorm:"type:varchar(100);not null"`
	ParentID  *string `gorm:"column:parent_id;type:varchar(255)"`
	OwnerID   string  `gorm:"column:owner_id;type:varchar(255);not null"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
