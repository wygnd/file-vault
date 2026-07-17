import { IsNotEmpty, IsUUID } from 'class-validator';

export class FileDeleteParamRequestDTO {
  @IsNotEmpty()
  @IsUUID()
  file_id: string;
}
