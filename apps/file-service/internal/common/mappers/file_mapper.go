package mappers

import (
	gen "github.com/wygnd/file-vault/file-service/gen/file"
	"github.com/wygnd/file-vault/file-service/internal/common/dto"
	"github.com/wygnd/file-vault/file-service/internal/common/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToFileResponseDTO(record *models.File) *dto.FileResponseDTO {
	return &dto.FileResponseDTO{
		ID:        record.ID,
		Name:      record.Name,
		MimeType:  record.MimeType,
		Size:      record.Size,
		CreatedAt: record.CreatedAt,
	}
}

func ToGrpcFileResponse(dto *dto.FileResponseDTO) *gen.FileResponse {
	return &gen.FileResponse{
		Id:        dto.ID,
		Name:      dto.Name,
		MimeType:  dto.MimeType,
		Size:      dto.Size,
		CreatedAt: timestamppb.New(dto.CreatedAt),
	}
}

func ToGrpcFileDetailResponse(dto *dto.FileDetailResponseDTO) *gen.GetByIdResponse {
	return &gen.GetByIdResponse{
		Id:        dto.ID,
		Name:      dto.Name,
		MimeType:  dto.MimeType,
		Size:      dto.Size,
		Url:       dto.URL,
		CreatedAt: timestamppb.New(dto.CreatedAt),
	}
}
