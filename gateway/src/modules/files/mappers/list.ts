import { fromGrpcFileToResponse } from '@modules/files/mappers/file';
import { FileListResponseDTO } from '@modules/files/dto';
import { ListByFolderIdResponse } from '@generated/file/file';

export const fromGrpcFileListResponse = (
  response: ListByFolderIdResponse,
): FileListResponseDTO => {
  return {
    items: response.result.map((r) => fromGrpcFileToResponse(r)),
    next_cursor: response.nextCursor,
  };
};
