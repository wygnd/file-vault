import {
  ArgumentsHost,
  Catch,
  ExceptionFilter,
  HttpException,
  HttpStatus,
} from '@nestjs/common';
import { FastifyReply } from 'fastify';
import { ErrorCode } from '@generated/enums/v1/enums';
import { status as GrpcStatus } from '@grpc/grpc-js';

@Catch()
export class TransformErrorFilter implements ExceptionFilter {
  catch(exception: unknown, host: ArgumentsHost) {
    const context = host.switchToHttp();
    const response = context.getResponse() as FastifyReply;

    const { status, errCode, errDetail } = this.resolveException(exception);

    response.status(status).send({
      ok: false,
      err_code: errCode,
      err_detail: errDetail,
      timestamp: new Date().toISOString(),
    });
  }

  private resolveException(exception: unknown): {
    status: number;
    errCode: string;
    errDetail: string;
  } {
    if (this.isGrpcError(exception)) {
      return {
        status: this.mapGrpcStatusToHttpStatus(exception.code),
        errCode: this.mapGrpcStatusToCode(exception.code),
        errDetail: exception.details,
      };
    }

    if (exception instanceof HttpException) {
      const response = exception.getResponse() as FastifyReply;
      const detail =
        typeof response === 'string'
          ? response
          : ((response as any).message ?? exception.message);

      return {
        status: exception.getStatus(),
        errCode: this.mapHttpStatusToCode(exception.getStatus()),
        errDetail: Array.isArray(detail) ? detail.join(', ') : detail,
      };
    }

    return {
      status: HttpStatus.INTERNAL_SERVER_ERROR,
      errCode: ErrorCode.INTERNAL_ERROR,
      errDetail: 'Internal Server Error',
    };
  }

  private isGrpcError(
    exception: unknown,
  ): exception is { code: number; details: string } {
    return (
      typeof exception === 'object' &&
      exception !== null &&
      'code' in exception &&
      'details' in exception
    );
  }

  private mapHttpStatusToCode(status: HttpStatus): ErrorCode {
    switch (status) {
      case HttpStatus.BAD_REQUEST:
        return ErrorCode.VALIDATION_ERROR;
      case HttpStatus.UNAUTHORIZED:
        return ErrorCode.UNAUTHORIZED;
      case HttpStatus.FORBIDDEN:
        return ErrorCode.FORBIDDEN;
      case HttpStatus.NOT_FOUND:
        return ErrorCode.NOT_FOUND;
      default:
        return ErrorCode.INTERNAL_ERROR;
    }
  }

  private mapGrpcStatusToHttpStatus(status: GrpcStatus): HttpStatus {
    switch (status) {
      case GrpcStatus.NOT_FOUND:
        return HttpStatus.NOT_FOUND;
      case GrpcStatus.INVALID_ARGUMENT:
        return HttpStatus.BAD_REQUEST;
      case GrpcStatus.PERMISSION_DENIED:
        return HttpStatus.FORBIDDEN;
      case GrpcStatus.UNAUTHENTICATED:
        return HttpStatus.UNAUTHORIZED;
      case GrpcStatus.ALREADY_EXISTS:
        return HttpStatus.CONFLICT;
      default:
        return HttpStatus.INTERNAL_SERVER_ERROR;
    }
  }

  private mapGrpcStatusToCode(status: GrpcStatus): ErrorCode {
    switch (status) {
      case GrpcStatus.NOT_FOUND:
        return ErrorCode.NOT_FOUND;
      case GrpcStatus.INVALID_ARGUMENT:
        return ErrorCode.VALIDATION_ERROR;
      case GrpcStatus.PERMISSION_DENIED:
        return ErrorCode.FORBIDDEN;
      case GrpcStatus.UNAUTHENTICATED:
        return ErrorCode.UNAUTHORIZED;
      case GrpcStatus.ALREADY_EXISTS:
        return ErrorCode.ALREADY_EXISTS;
      default:
        return ErrorCode.INTERNAL_ERROR;
    }
  }
}
