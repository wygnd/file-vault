import {
  BadRequestException,
  HttpException,
  Injectable,
  InternalServerErrorException,
  Logger,
} from '@nestjs/common';
import { MultipartFile } from '@fastify/multipart';
import { FileGrpcProvider } from '@modules/files/providers';

@Injectable()
export class FileService {
  private readonly logger = new Logger(FileService.name);

  constructor(private readonly fileGrpcProvider: FileGrpcProvider) {}

  public async handleUpload(file: MultipartFile) {
    try {
      const fileBuffer = await file.toBuffer();

      if (fileBuffer.length == 0) {
        throw new BadRequestException('No uploaded file');
      }

      return await this.fileGrpcProvider.uploadFile({
        fileName: file.filename,
        size: fileBuffer.length,
        mimeType: file.mimetype,
        ownerId: '123',
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
}
