import { Controller, Post } from '@nestjs/common';
import { ApiOperation } from '@nestjs/swagger';

@Controller({
  version: '1',
  path: 'files',
})
export class FileControllerV1 {
  constructor() {}

  @ApiOperation({ summary: 'Загрузить файл' })
  @Post('')
  public async upload() {}
}
