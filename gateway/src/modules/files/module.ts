import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { FILE_SERVICE } from '@modules/files/constants/constants';
import { FILE_PACKAGE_NAME } from '@generated/file/file';
import { join } from 'node:path';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: FILE_SERVICE,
        useFactory: () => ({
          name: FILE_SERVICE,
          transport: Transport.GRPC,
          options: {
            package: FILE_PACKAGE_NAME,
            protoPath: join(
              process.cwd(),
              'shared',
              'proto',
              'file',
              'file.proto',
            ),
          },
        }),
      },
    ]),
  ],
})
export class FilesModule {}
