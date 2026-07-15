import { FastifyAdapter } from '@nestjs/platform-fastify';
import multipart from '@fastify/multipart';

export const setupAppPlugins = async (adapter: FastifyAdapter) => {
  await adapter.register(multipart, {
    limits: {
      fileSize: 10 * 1024 * 1024, // 10 mb
    },
  });
};
