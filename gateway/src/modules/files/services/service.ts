import {
  BadRequestException,
  HttpException,
  Injectable,
  InternalServerErrorException,
  Logger,
} from '@nestjs/common';
import { MultipartFile } from '@fastify/multipart';
import { FileGrpcProvider } from '@modules/files/providers';
import { FileListQueryRequestDTO } from '@modules/files/dto';

@Injectable()
export class FileService {
  private readonly logger = new Logger(FileService.name);

  constructor(private readonly fileGrpcProvider: FileGrpcProvider) {}

  /**
   * Получает файл и отправляет через провайдер на загрузку
   * @param file
   */
  public async upload(file: MultipartFile) {
    try {
      const fileBuffer = await file.toBuffer();

      if (fileBuffer.length == 0) {
        throw new BadRequestException('No uploaded file');
      }

      return await this.fileGrpcProvider.uploadFile({
        fileName: file.filename,
        size: fileBuffer.length,
        mimeType: file.mimetype,
        ownerId: '123', // fixme
        data: fileBuffer,
      });
    } catch (error) {
      if (error instanceof HttpException) {
        throw error;
      }

      this.logger.error(error);
      throw new InternalServerErrorException(
        'Произошла непредвиденная ошибка на сервере',
      );
    }
  }

  /**
   * Получает список файлов
   *
   * Возвращает список файлов в директории
   * @param query
   */
  public async getFileList(query: FileListQueryRequestDTO) {
    try {
      const result = await this.fileGrpcProvider.getFiles(
        query.folder_id,
        query.cursor,
        query.limit,
      );

      return {
        items: result.result,
        next_cursor: result.nextCursor,
      };
    } catch (error) {
      if (error instanceof HttpException) {
        throw error;
      }

      this.logger.error(error);
      throw new InternalServerErrorException(
        'Произошла непредвиденная ошибка на сервере',
      );
    }
  }

  /**
   * Удаляет файл
   * @param fileId
   */
  public async deleteFile(fileId: string) {
    try {
      await this.fileGrpcProvider.delete(fileId);

      return true;
    } catch (error) {
      if (error instanceof HttpException) {
        throw error;
      }

      this.logger.error(error);
      throw new InternalServerErrorException(
        'Произошла непредвиденная ошибка на сервере',
      );
    }
  }
}
