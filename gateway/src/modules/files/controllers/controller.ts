import {
  BadRequestException,
  Controller,
  Get,
  Post,
  Req,
} from '@nestjs/common';
import { ApiOperation, ApiTags } from '@nestjs/swagger';
import { type FastifyRequest } from 'fastify';
import { FileService } from '@modules/files/services/service';

@ApiTags('Работа с файлами')
@Controller({
  version: '1',
  path: 'files',
})
export class FileControllerV1 {
  constructor(private readonly fileService: FileService) {}

  @ApiOperation({ summary: 'Загрузить файл' })
  @Post('upload')
  public async uploadFile(@Req() request: FastifyRequest) {
    if (!request.isMultipart()) {
      throw new BadRequestException('Multipart form data expected');
    }

    const file = await request.file();

    if (!file) {
      throw new BadRequestException('File upload failed');
    }

    return this.fileService.handleUpload(file);
  }

  @ApiOperation({ summary: 'Получить список файлов' })
  @Get()
  public async getFiles() {
    // todo
  }
}
