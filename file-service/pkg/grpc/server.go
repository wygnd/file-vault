package grpcFileService

import (
	"context"

	gen "github.com/wygnd/file-vault/file-service/gen/file"
	"github.com/wygnd/file-vault/file-service/internal/common/dto"
	"github.com/wygnd/file-vault/file-service/internal/common/mappers"
	file_service "github.com/wygnd/file-vault/file-service/internal/common/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FileGrpcService struct {
	gen.UnimplementedFileServiceServer
	service file_service.FileService
}

func NewFileGrpcService(service file_service.FileService) *FileGrpcService {
	return &FileGrpcService{service: service}
}

func (server *FileGrpcService) Upload(ctx context.Context, request *gen.UploadRequest) (*gen.FileResponse, error) {
	result, err := server.service.Upload(dto.UploadFile{
		Name:     request.FileName,
		MimeType: request.MimeType,
		Size:     request.Size,
		Data:     request.Data,
	})

	if err != nil {
		return nil, err
	}

	return mappers.ToGrpcFileResponse(result), nil
}

func (server *FileGrpcService) GetById(ctx context.Context, request *gen.GetByIdRequest) (*gen.GetByIdResponse, error) {
	result, err := server.service.GetByID(request.Id)

	if err != nil {
		return nil, err
	}

	return &gen.GetByIdResponse{
		Url: result,
	}, nil
}

func (server *FileGrpcService) Delete(ctx context.Context, request *gen.DeleteRequest) (*emptypb.Empty, error) {
	err := server.service.Delete(request.Id)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (server *FileGrpcService) ListByFolderId(ctx context.Context, request *gen.ListByFolderIdRequest) (*gen.ListByFolderIdResponse, error) {

	result, err := server.service.ListByFolderID(request.FolderId)

	if err != nil {
		return nil, err
	}

	resultList := make([]*gen.FileResponse, 0, len(result))

	for _, res := range result {
		resultList = append(resultList, mappers.ToGrpcFileResponse(res))
	}

	return &gen.ListByFolderIdResponse{
		Result: resultList,
	}, nil
}
