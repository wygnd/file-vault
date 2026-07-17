package mappers

import (
	gen "github.com/wygnd/file-vault/file-service/gen/file"
	"github.com/wygnd/file-vault/file-service/internal/common/dto"
	"github.com/wygnd/file-vault/file-service/internal/common/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToFolderResponseDTO(record *models.Folder) *dto.FolderResponseDTO {
	return &dto.FolderResponseDTO{
		ID:        record.ID,
		Name:      record.Name,
		ParentID:  record.ParentID,
		OwnerID:   record.OwnerID,
		CreatedAt: record.CreatedAt,
	}
}

func ToGrpcFolderResponseDTO(dto *dto.FolderResponseDTO) *gen.FolderResponse {
	return &gen.FolderResponse{
		Id:        dto.ID,
		Name:      dto.Name,
		ParentId:  dto.ParentID,
		CreatedAt: timestamppb.New(dto.CreatedAt),
	}
}

func FromGrpcUpdateFolderRequestDTO(request *gen.UpdateFolderRequest) *dto.UpdateFolderDTO {
	return &dto.UpdateFolderDTO{
		Name: request.Name,
	}
}

func ToGrpcFolderResponseDTOList(folders []*dto.FolderResponseDTO) []*gen.FolderResponse {
	dtoFolderList := make([]*gen.FolderResponse, 0, len(folders))

	for _, folder := range folders {
		dtoFolderList = append(dtoFolderList, &gen.FolderResponse{
			Id:        folder.ID,
			Name:      folder.Name,
			ParentId:  folder.ParentID,
			CreatedAt: timestamppb.New(folder.CreatedAt),
		})
	}

	return dtoFolderList
}

func ToGrpcFileResponseDTOList(files []*dto.FileResponseDTO) []*gen.FileResponse {
	dtoFileList := make([]*gen.FileResponse, 0, len(files))

	for _, file := range files {
		dtoFileList = append(dtoFileList, &gen.FileResponse{
			Id:        file.ID,
			Name:      file.Name,
			MimeType:  file.MimeType,
			Size:      file.Size,
			CreatedAt: timestamppb.New(file.CreatedAt),
		})
	}

	return dtoFileList
}

func ToGrpcGetFolderChildrenResponseDTO(dto *dto.FolderChildrenResponseDTO) *gen.GetFolderChildrenResponse {
	return &gen.GetFolderChildrenResponse{
		Folders: ToGrpcFolderResponseDTOList(dto.Folders),
		Files:   ToGrpcFileResponseDTOList(dto.Files),
	}
}
