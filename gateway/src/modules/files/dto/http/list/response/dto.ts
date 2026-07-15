import { FileResponseDTO } from '@modules/files/dto';

export class FileListResponseDTO {
  items: FileResponseDTO[];
  next_cursor?: string;
}
