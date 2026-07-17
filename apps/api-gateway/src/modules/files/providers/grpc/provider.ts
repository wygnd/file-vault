import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import { FILE_SERVICE_NAME, FileServiceClient } from '@generated/file/file';
import { type ClientGrpc } from '@nestjs/microservices';
import { FILE_SERVICE } from '@modules/files/constants/constants';
import {
  FileDetailResponseDTO,
  FileResponseDTO,
  UploadFileGrpcDTO,
} from '@modules/files/dto';
import { firstValueFrom } from 'rxjs';
import { fromGrpcFileListResponse } from '@modules/files/mappers/list';
import {
  fromGrpcFileToDetailResponse,
  fromGrpcFileToResponse,
} from '@modules/files/mappers/file';

@Injectable()
export class FileGrpcProvider implements OnModuleInit {
  private readonly logger = new Logger(FileGrpcProvider.name);
  private fileClient: FileServiceClient;

  constructor(
    @Inject(FILE_SERVICE)
    private readonly grpcClient: ClientGrpc,
  ) {}

  public onModuleInit(): void {
    this.fileClient = this.grpcClient.getService(FILE_SERVICE_NAME);
  }

  /**
   * Отправляет запрос на загрузку файла
   * @param file
   */
  public async uploadFile(file: UploadFileGrpcDTO): Promise<FileResponseDTO> {
    try {
      const result = await firstValueFrom(this.fileClient.upload(file));

      return fromGrpcFileToResponse(result);
    } catch (error) {
      this.logger.error(error);
      throw error;
    }
  }

  /**
   * Получить список файлов
   * @param folderId
   * @param cursor
   * @param limit
   */
  public async getFiles(
    folderId: string = '',
    cursor?: string,
    limit: number = 25,
  ) {
    try {
      const result = await firstValueFrom(
        this.fileClient.listByFolderId({
          folderId: folderId,
          cursor: cursor,
          limit: limit,
        }),
      );
      return fromGrpcFileListResponse(result);
    } catch (error) {
      this.logger.error(error);
      throw error;
    }
  }

  /**
   * Получить файл по ID
   * @param fileId
   */
  public async getById(fileId: string): Promise<FileDetailResponseDTO> {
    try {
      const result = await firstValueFrom(
        this.fileClient.getById({
          id: fileId,
        }),
      );

      return fromGrpcFileToDetailResponse(result);
    } catch (error) {
      this.logger.error(error);
      throw error;
    }
  }

  /**
   * Удалить файл по ID
   * @param fileId
   */
  public async delete(fileId: string) {
    try {
      return await firstValueFrom(
        this.fileClient.delete({
          id: fileId,
        }),
      );
    } catch (error) {
      this.logger.error(error);
      throw error;
    }
  }
}
