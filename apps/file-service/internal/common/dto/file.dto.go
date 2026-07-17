package dto

import "time"

type UploadFile struct {
	Name     string
	MimeType string
	Size     int64
	Data     []byte
	OwnerId  string
	FolderId string
}

type FileResponseDTO struct {
	ID        string
	Name      string
	MimeType  string
	Size      int64
	CreatedAt time.Time
}

type FileDetailResponseDTO struct {
	FileResponseDTO
	URL string
}
