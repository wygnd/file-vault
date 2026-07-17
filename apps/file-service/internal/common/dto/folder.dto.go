package dto

import "time"

type UpdateFolderDTO struct {
	Name string
}

type CreateFolderDTO struct {
	Name     string
	ParentID *string
	OwnerID  string
}

type FolderResponseDTO struct {
	ID        string
	Name      string
	ParentID  *string
	OwnerID   string
	CreatedAt time.Time
}

type FolderChildrenResponseDTO struct {
	Folders []*FolderResponseDTO
	Files   []*FileResponseDTO
}
