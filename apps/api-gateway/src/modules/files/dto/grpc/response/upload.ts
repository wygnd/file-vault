import { FileResponse } from '@generated/file/file';
import { Timestamp } from '@generated/google/protobuf/timestamp';

export class UploadFileGrpcResponseDTO implements FileResponse {
  id: string;
  mimeType: string;
  name: string;
  size: number;
  createdAt: Timestamp | undefined;
  folderId?: string;
}
