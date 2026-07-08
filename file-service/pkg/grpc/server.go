package grpcFileService

import (
	"context"

	gen "github.com/wygnd/file-vault/file-service/gen/file"
	"github.com/wygnd/file-vault/file-service/internal/common/dto"
	"github.com/wygnd/file-vault/file-service/internal/common/mappers"
	"github.com/wygnd/file-vault/file-service/internal/common/service"
	"google.golang.org/protobuf/types/known/emptypb"
)

type FileGrpcService struct {
	gen.UnimplementedFileServiceServer
	fileService   service.FileService
	folderService service.FolderService
}

/* ================ Файлы ================ */

func NewFileGrpcService(fileService service.FileService, folderService service.FolderService) *FileGrpcService {
	return &FileGrpcService{fileService: fileService, folderService: folderService}
}

func (server *FileGrpcService) Upload(ctx context.Context, request *gen.UploadRequest) (*gen.FileResponse, error) {
	result, err := server.fileService.Upload(dto.UploadFile{
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
	result, err := server.fileService.GetByID(request.Id)

	if err != nil {
		return nil, err
	}

	return mappers.ToGrpcFileDetailResponse(result), nil
}

func (server *FileGrpcService) Delete(ctx context.Context, request *gen.DeleteRequest) (*emptypb.Empty, error) {
	err := server.fileService.Delete(request.Id)

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (server *FileGrpcService) ListByFolderId(ctx context.Context, request *gen.ListByFolderIdRequest) (*gen.ListByFolderIdResponse, error) {

	result, err := server.fileService.ListByFolderID(request.FolderId)

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

/* ================ Папки ================ */

// CreateFolder создает папку
func (server *FileGrpcService) CreateFolder(ctx context.Context, request *gen.CreateFolderRequest) (*gen.FolderResponse, error) {
	result, err := server.folderService.Create(dto.CreateFolderDTO{
		Name:     request.Name,
		ParentID: request.ParentId,
		OwnerID:  request.OwnerId,
	})

	if err != nil {
		return nil, err
	}

	return mappers.ToGrpcFolderResponseDTO(result), nil
}

// UpdateFolder обновляет папку
func (server *FileGrpcService) UpdateFolder(crt context.Context, request *gen.UpdateFolderRequest) (*gen.FolderResponse, error) {
	result, err := server.folderService.Update(request.Id, mappers.FromGrpcUpdateFolderRequestDTO(request))

	if err != nil {
		return nil, err
	}

	return mappers.ToGrpcFolderResponseDTO(result), nil
}

// DeleteFolder удаляет папку
func (server *FileGrpcService) DeleteFolder(ctx context.Context, request *gen.DeleteFolderRequest) (*emptypb.Empty, error) {
	err := server.folderService.Delete(request.Id)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

// GetFolderChildren получает список файлов и папок
func (server *FileGrpcService) GetFolderChildren(ctx context.Context, request *gen.GetFolderChildrenRequest) (*gen.GetFolderChildrenResponse, error) {
	result, err := server.folderService.GetChildren(request.OwnerId, request.ParentId)

	if err != nil {
		return nil, err
	}

	return mappers.ToGrpcGetFolderChildrenResponseDTO(result), nil
}
