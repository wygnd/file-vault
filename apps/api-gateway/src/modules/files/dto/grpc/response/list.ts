import { FileResponse } from '@generated/file/file';

export class FileGrpcListResponseDTO {
  items: FileResponse[];
  next_cursor?: string;
}
