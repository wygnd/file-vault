import { IsNotEmpty, IsString, IsUUID } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class FileGetParamsDTO {
  @ApiProperty({
    type: String,
    description: 'ID файла',
    required: true,
  })
  @IsNotEmpty()
  @IsString()
  @IsUUID()
  file_id: string;
}
