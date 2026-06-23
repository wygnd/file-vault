import { Inject, Injectable, Logger, OnModuleInit } from '@nestjs/common';
import { FILE_SERVICE_NAME, FileServiceClient } from '@generated/file/file';
import { type ClientGrpc } from '@nestjs/microservices';
import { FILE_SERVICE } from '@modules/files/constants/constants';
import { UploadFileGrpcDTO } from '@modules/files/dto';
import { UploadFileGrpcResponseDTO } from '@modules/files/dto/grpc/response/upload';
import { firstValueFrom } from 'rxjs';

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
  public async uploadFile(
    file: UploadFileGrpcDTO,
  ): Promise<UploadFileGrpcResponseDTO> {
    try {
      return await firstValueFrom(this.fileClient.upload(file));
    } catch (error) {
      this.logger.error(error);
      throw error;
    }
  }

  // todo
  public async getFiles(folderId: string = '') {
    try {
      return await firstValueFrom(
        this.fileClient.listByFolderId({
          folderId: folderId,
        }),
      );
    } catch (error) {
      this.logger.error(error);
      throw error;
    }
  }
}
