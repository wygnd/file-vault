# 📁 FileVault — Personal File Storage Service

Учебный проект: самописный файловый сервис по типу iCloud.  
Цель — изучить **Angular**, **Go** и **NestJS** на реальном проекте.

---

## Архитектура

```
┌─────────────────────────────────────────────────────┐
│                   Клиенты                           │
│         Angular SPA        Telegram Bot             │ 
│                             (in future)             │
└────────────────┬───────────────────┬────────────────┘
                 │                   │
                 ▼                   ▼
┌─────────────────────────────────────────────────────┐
│              NestJS API Gateway                     │
│   - Аутентификация (JWT)                            │
│   - Авторизация (роли)                              │
│   - Маршрутизация запросов                          │
└────────────────────────────┬────────────────────────┘
                             │ gRPC / REST
                             ▼
┌─────────────────────────────────────────────────────┐
│              Go File Service                        │
│   - Загрузка / скачивание файлов                    │
│   - Генерация превью (thumbnails)                   │
│   - Управление метаданными                          │
└──────────────┬──────────────────────┬───────────────┘
               │                      │
               ▼                      ▼
       ┌───────────────┐     ┌────────────────┐
       │  PostgreSQL   │     │     MinIO      │
       │  (метаданные) │     │  (файлы/фото)  │
       └───────────────┘     └────────────────┘
```

### Компоненты

| Компонент | Технология | Назначение |
|---|---|---|
| Frontend | Angular 17+ | SPA: просмотр файлов, загрузка, управление папками |
| API Gateway | NestJS + TypeScript | Аутентификация, роутинг, валидация |
| File Service | Go (Gin/Echo) | Работа с файлами, стриминг, превью |
| База данных | PostgreSQL | Пользователи, папки, метаданные файлов |
| Хранилище | MinIO | Бинарное хранение файлов (S3-совместимо) |
| Инфраструктура | Docker Compose | Локальный запуск всех сервисов |

---

## Схема базы данных

```sql
users
  id, email, password_hash, role, created_at

folders
  id, name, parent_id, owner_id, created_at   -- рекурсивная структура

files
  id, name, folder_id, owner_id,
  storage_key,   -- UUID-ключ в MinIO (не оригинальное имя!)
  mime_type, size, hash_sha256,
  created_at
```

---

## Структура репозитория

```
filevault/
├── frontend/          # Angular приложение
├── gateway/           # NestJS API Gateway
├── file-service/      # Go сервис
├── docker-compose.yml
└── README.md
```

---

## TODO

### 🔧 Инфраструктура
- [x] Создать `docker-compose.yml`
  - [x]  PostgreSQL
  - [x]  MinIO
- [x] Настроить `.env` файлы для каждого сервиса
- [x] Написать `Makefile`

---

### 🐹 Go File Service (`/file-service`)
- [x] Инициализировать проект (`go mod init`)
- [x] Подключить MinIO SDK (`minio-go`)
- [x] Подключить PostgreSQL (`pgx` или `sqlx`)
- [x] Сохранять метаданные файла в БД (UUID-ключ, mime, size, hash)
- [x] SHA-256 хэш при загрузке (дедупликация)
- [ ] Генерация thumbnails для изображений (`imaging` или `vips`)
- [ ] Unit-тесты для хендлеров
- [x] Dockerfile для сервиса

---

### 🏗️ NestJS API Gateway (`/gateway`)
- [x] Инициализировать проект (`nest new`)
- [ ] Guards для защиты роутов
- [ ] Модуль пользователей + роли
- [ ] Проксирование запросов в Go file service
- [ ] Модуль папок (CRUD)
- [ ] Валидация через `class-validator`
- [x] Swagger документация (`@nestjs/swagger`)
- [x] Dockerfile для сервиса

---

### 🅰️ Angular Frontend (`/frontend`)
- [ ] Инициализировать проект (`ng new`)
- [ ] Настроить Angular Router + Auth Guard
- [ ] HTTP Interceptor для JWT токенов
- [ ] Страница авторизации (login/register)
- [ ] Главная страница: файловый менеджер (список папок и файлов)
- [ ] Компонент загрузки файлов (drag & drop)
- [ ] Компонент создания папки
- [ ] Просмотр изображений (галерея/лайтбокс)
- [ ] Индикатор прогресса загрузки
- [ ] Адаптив под мобильные устройства

---

### 🤖 Telegram Bot (опционально)
- [ ] Создать бота через @BotFather
- [ ] Go или Node.js обёртка
- [ ] Команда `/upload` — отправить файл боту → сохраняется на сервер
- [ ] Команда `/list` — список последних загрузок
- [ ] Учесть лимит Telegram: файлы до 2 ГБ для ботов

---

## Запуск

```bash
# Поднять инфраструктуру
docker compose up -d

# Go сервис (dev)
cd file-service && go run ./cmd/main.go

# NestJS gateway (dev)
cd gateway && npm run start:dev

# Angular (dev)
cd frontend && ng serve
```

---

## Полезные ссылки

- [MinIO Go SDK](https://github.com/minio/minio-go)
- [NestJS Docs](https://docs.nestjs.com)
- [Angular Docs](https://angular.dev)
- [pgx — PostgreSQL driver for Go](https://github.com/jackc/pgx)
- [Gin Web Framework](https://gin-gonic.com)
