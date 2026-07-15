import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { FilesModule } from '@modules/files/module';
import { APP_INTERCEPTOR } from '@nestjs/core';
import { TransformSuccessResponseInterceptor } from '@shared/interceptors';

@Module({
  imports: [
    // System modules
    ConfigModule.forRoot({
      // envFilePath: join(process.cwd(), '..', '.env'), // fixme: not for dev
      isGlobal: true,
      cache: true,
    }),

    // Other modules
    FilesModule,
  ],
  controllers: [],
  providers: [
    {
      provide: APP_INTERCEPTOR,
      useClass: TransformSuccessResponseInterceptor,
    },
  ],
})
export class AppModule {}
