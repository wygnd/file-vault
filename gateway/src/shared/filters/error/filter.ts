import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpStatus,
} from '@nestjs/common';
import { FastifyReply } from 'fastify';

@Catch()
export class TransformErrorFilter implements ExceptionFilter {
  catch(exception: unknown, host: ArgumentsHost) {
    const context = host.switchToHttp(),
      response = context.getResponse() as FastifyReply;

    let status = HttpStatus.INTERNAL_SERVER_ERROR;

    response.status(status).send({
      ok: false,
      err_code: '',
      err_detail: '',
      timestamp: new Date().toISOString(),
    });
  }
}
