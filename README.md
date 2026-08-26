# Golang ToDo App

REST API для управления пользователями, задачами и статистикой выполнения задач.
Проект написан на Go и использует PostgreSQL, стандартный `net/http`, `pgx`,
`zap`, `envconfig` и Swagger через `swaggo`.

## Возможности

- CRUD для пользователей.
- CRUD для задач.
- Фильтрация задач по автору, пагинация списков.
- Частичное обновление через PATCH с three-state logic.
- Статистика по созданным и завершенным задачам.
- Swagger UI.
- Простая HTML-страница из каталога `public`.

## Быстрый старт

1. Создайте `.env` на основе `.env.example` и заполните параметры PostgreSQL.
2. Поднимите окружение:

```sh
make env-up
make migrate-up
make todoapp-run
```

После запуска приложение доступно на `http://127.0.0.1:5050`.

Swagger UI:

```text
http://127.0.0.1:5050/swagger/
```

## Docker-запуск

```sh
make todoapp-deploy
```

Остановка приложения:

```sh
make todoapp-undeploy
```

Остановка окружения PostgreSQL:

```sh
make env-down
```

## Структура проекта

```text
cmd/todoapp/        Точка входа и сборка зависимостей приложения
internal/core/      Общая инфраструктура: config, logger, HTTP, PostgreSQL pool
internal/features/  Бизнес-фичи: users, tasks, statistics, web
migrations/         SQL-миграции PostgreSQL
docs/               Swagger-артефакты и проектная документация
public/             Статическая HTML-страница
```

## Документация

- [Архитектура](docs/architecture.md)
- [Конфигурация](docs/configuration.md)
- [Локальная разработка](docs/local-development.md)
- [API](docs/api.md)
- [База данных](docs/database.md)
- [Эксплуатация](docs/operations.md)
- [Тестирование](docs/testing.md)
- [Правила участия](docs/contributing.md)
- [Правила комментариев](docs/commenting-guidelines.md)

## Swagger

Swagger-спецификация генерируется в `docs/swagger.json`, `docs/swagger.yaml`
и `docs/docs.go`.

```sh
make swagger-gen
```
