export class FileResponseDTO {
  id: string;
  mimeType: string;
  name: string;
  size: number;
  createdAt: number;
  folderId?: string;
}

export class FileDetailResponseDTO extends FileResponseDTO {
  url: string;
}
