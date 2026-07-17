import { IsInt, IsOptional, IsString } from 'class-validator';
import { Type } from 'class-transformer';
import { ApiProperty } from '@nestjs/swagger';

export class FileListQueryRequestDTO {
  @ApiProperty({
    type: Number,
    description: 'Количество файлов на странице',
    required: false,
    default: 25,
    example: 25,
  })
  @IsOptional()
  @Type(() => Number)
  @IsInt({ message: 'Limit must be a positive integer' })
  limit: number = 25;

  @ApiProperty({
    type: String,
    description: 'ID файла, после которого нужно получать файлы',
    required: false,
    example: '123',
  })
  @IsOptional()
  @IsString()
  cursor?: string;

  @ApiProperty({
    type: String,
    description: 'ID папки для получения дочерних элементов',
    required: false,
    example: '123',
  })
  @IsOptional()
  @IsString()
  folder_id?: string;
}
