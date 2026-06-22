FROM node:24-alpine AS build
LABEL authors="wygnd"

WORKDIR /app

COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./

RUN corepack enable
RUN pnpm install --frozen-lockfile

COPY . .

RUN pnpm run build


FROM node:24-alpine AS production

WORKDIR /app

RUN corepack enable

COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./

RUN pnpm install --frozen-lockfile --prod

COPY --from=build /app/dist ./dist
COPY ./gen ./dist/gen

EXPOSE 5000

CMD ["node", "dist/src/main.js"]
