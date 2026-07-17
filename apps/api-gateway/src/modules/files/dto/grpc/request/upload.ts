import { UploadRequest } from '@generated/file/file';

export class UploadFileGrpcDTO implements UploadRequest {
  fileName: string;
  mimeType: string;
  size: number;
  data: Uint8Array;
  ownerId: string;
  folderId?: string;
}
