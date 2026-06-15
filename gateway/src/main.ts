import { NestFactory } from '@nestjs/core';
import { AppModule } from '@modules/module';
import {
  FastifyAdapter,
  NestFastifyApplication,
} from '@nestjs/platform-fastify';
import { setupAppVersioning } from './common/versioning';
import { ConfigService } from '@nestjs/config';
import { Logger } from '@nestjs/common';
import { setupAppCors } from './common/cors';
import { setupAppPrefix } from './common/preffix';
import { setupAppFilters } from './common/filters';
import { setupAppPipes } from './common/pipes';
import { setupAppDocs } from './common/documentation';

async function bootstrap() {
  const app = await NestFactory.create<NestFastifyApplication>(
    AppModule,
    new FastifyAdapter(),
  );
  const config = app.get(ConfigService);
  const logger = new Logger('Application');

  // Добавляем префикс к приложению
  setupAppPrefix(app);

  // Добавляем версионирование
  setupAppVersioning(app);

  // Добавляем CORS
  setupAppCors(app);

  // Добавляем фильтры
  setupAppFilters(app);

  // Добавляем pipes
  setupAppPipes(app);

  // Добавляем документацию
  setupAppDocs(app);

  const PORT = config.get<number>('API_GATEWAY_PORT') ?? 3000;

  await app.listen(PORT, '0.0.0.0');
  logger.log(`Server started http:/0.0.0.0:${PORT}`);
}

bootstrap();
