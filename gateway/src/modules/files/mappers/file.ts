import { FileResponse, GetByIdResponse } from '@generated/file/file';
import { FileDetailResponseDTO, FileResponseDTO } from '@modules/files/dto';
import Long from 'long';

export const fromGrpcFileToResponse = (file: FileResponse): FileResponseDTO => {
  let createdAt: number = 0;

  if (file.createdAt) {
    createdAt = new Date(file.createdAt.seconds * 1000).getTime();
  }

  return {
    id: file.id,
    name: file.name,
    size: Long.fromValue(file.size).toNumber(),
    createdAt: createdAt,
    mimeType: file.mimeType,
  };
};

export const fromGrpcFileToDetailResponse = (
  file: GetByIdResponse,
): FileDetailResponseDTO => {
  let createdAt: number = 0;

  if (file.createdAt) {
    createdAt = new Date(file.createdAt.seconds * 1000).getTime();
  }

  return {
    id: file.id,
    name: file.name,
    size: Long.fromValue(file.size).toNumber(),
    createdAt: createdAt,
    mimeType: file.mimeType,
    url: file.url,
  };
};
