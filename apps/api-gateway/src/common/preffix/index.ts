import {NestFastifyApplication} from "@nestjs/platform-fastify";
import {RequestMethod} from "@nestjs/common";

export const setupAppPrefix = (app: NestFastifyApplication): void => {
	app.setGlobalPrefix('api', {exclude: [{path: 'docs', method: RequestMethod.GET}]})
}