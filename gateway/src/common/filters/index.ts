import { NestFastifyApplication } from '@nestjs/platform-fastify';
import { TransformErrorFilter } from '@shared/filters/error/filter';

export const setupAppFilters = (app: NestFastifyApplication): void => {
  app.useGlobalFilters(new TransformErrorFilter());
};
