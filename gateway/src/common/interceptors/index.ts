import { NestFastifyApplication } from '@nestjs/platform-fastify';
import { TransformSuccessResponseInterceptor } from '@shared/interceptors';

export const setupAppInterceptors = (app: NestFastifyApplication) => {
  app.useGlobalInterceptors(new TransformSuccessResponseInterceptor());
};
