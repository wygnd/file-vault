import { Module } from '@nestjs/common';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: '',
        useFactory: (config: ConfigService) => ({
          name: '',
          transport: Transport.GRPC,
          options: {
            package: '',
            protoPath: '',
          },
        }),
        inject: [ConfigService],
      },
    ]),
  ],
})
export class FilesModule {}
