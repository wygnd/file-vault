import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { FILE_SERVICE } from '@modules/files/constants/constants';
import { FILE_PACKAGE_NAME } from '@generated/file/file';
import { join } from 'node:path';
import { FileControllerV1 } from '@modules/files/controllers/controller';
import { FileService } from '@modules/files/services/service';
import { FileGrpcProvider } from '@modules/files/providers';
import { ConfigService } from '@nestjs/config';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: FILE_SERVICE,
        useFactory: (config: ConfigService) => ({
          name: FILE_SERVICE,
          transport: Transport.GRPC,
          options: {
            url: config.getOrThrow('API_GATEWAY_FILE_SERVICE_URL'),
            package: FILE_PACKAGE_NAME,
            protoPath: join(
              process.cwd(),
              '..',
              'shared',
              'proto',
              'file',
              'file.proto',
            ),
          },
        }),
        inject: [ConfigService],
      },
    ]),
  ],
  controllers: [FileControllerV1],
  providers: [FileGrpcProvider, FileService],
})
export class FilesModule {}
