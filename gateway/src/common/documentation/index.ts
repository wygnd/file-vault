import { NestFastifyApplication } from '@nestjs/platform-fastify';
import { ConfigService } from '@nestjs/config';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { apiReference } from '@scalar/nestjs-api-reference';

export const setupAppDocs = (app: NestFastifyApplication): void => {
  const config = app.get(ConfigService);
  const title =
    config.get('app.api.title', { infer: true }) ?? 'Api Mini Gateway';
  const description = config.get('app.api.description', { infer: true }) ?? '';

  const swaggerConfig = new DocumentBuilder()
    .setTitle(title)
    .setVersion('1.0')
    .setDescription(description)
    .addBearerAuth()
    .build();

  const document = SwaggerModule.createDocument(app, swaggerConfig);
  const handler = apiReference({
    withFastify: true,
    content: document,
    theme: 'fastify',
    showDeveloperTools: 'never',
    agent: {
      disabled: true,
    },
    mcp: {
      disabled: true,
    },
    pageTitle: title,
  });

  const fastify = app.getHttpAdapter().getInstance();

  fastify.get('/docs', async (req, reply) => {
    (handler as (req: unknown, res: NodeJS.WritableStream) => void)(
      req.raw,
      reply.raw,
    );
  });

  fastify.get('/docs/*', async (req, reply) => {
    (handler as (req: unknown, res: NodeJS.WritableStream) => void)(
      req.raw,
      reply.raw,
    );
  });
};
