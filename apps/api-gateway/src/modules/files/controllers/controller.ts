import {
  BadRequestException,
  Controller,
  Delete,
  Get,
  Param,
  Post,
  Query,
  Req,
} from '@nestjs/common';
import {
  ApiConsumes,
  ApiOkResponse,
  ApiOperation,
  ApiParam,
  ApiTags,
} from '@nestjs/swagger';
import { type FastifyRequest } from 'fastify';
import { FileService } from '@modules/files/services/service';
import {
  FileDeleteParamRequestDTO,
  FileGetParamsDTO,
  FileListQueryRequestDTO,
} from '@modules/files/dto';

@ApiTags('Работа с файлами')
@Controller({
  version: '1',
  path: 'files',
})
export class FileControllerV1 {
  constructor(private readonly fileService: FileService) {}

  @ApiOperation({ summary: 'Загрузить файл' })
  @ApiConsumes('multipart/form-data')
  @Post()
  public async uploadFile(@Req() request: FastifyRequest) {
    if (!request.isMultipart()) {
      throw new BadRequestException('Multipart form data expected');
    }

    const file = await request.file();

    if (!file) {
      throw new BadRequestException('File upload failed');
    }

    return this.fileService.upload(file);
  }

  @ApiOperation({ summary: 'Получить список файлов' })
  @Get()
  public async getFiles(@Query() query: FileListQueryRequestDTO) {
    return this.fileService.getFileList(query);
  }

  @ApiOperation({ summary: 'Удалить файл' })
  @ApiParam({
    name: 'file_id',
    description: 'ID файла',
    required: true,
    example: '123',
  })
  @ApiOkResponse({
    type: Boolean,
    description: 'Успех',
    example: true,
  })
  @Delete(':file_id')
  async deleteFile(@Param() params: FileDeleteParamRequestDTO) {
    return this.fileService.deleteFile(params.file_id);
  }

  @ApiOperation({ summary: 'Получить информацию о файле по ID' })
  @Get(':file_id')
  async getFile(@Param() params: FileGetParamsDTO) {
    return this.fileService.getFileById(params.file_id);
  }
}
