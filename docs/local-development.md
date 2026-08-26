# Локальная разработка

Основные сценарии разработки описаны в `Makefile`.

## Подготовка

1. Скопируйте `.env.example` в `.env`.
2. Заполните `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`.
3. Убедитесь, что установлен Docker.

## Запуск базы

```sh
make env-up
```

Команда поднимает сервис `todoapp-postgres` из `docker-compose.yaml`.

## Миграции

Применить миграции:

```sh
make migrate-up
```

Откатить миграции:

```sh
make migrate-down
```

Создать новую миграцию:

```sh
make migrate-create seq=add_some_table
```

## Port forwarding для локального Go-запуска

```sh
make env-port-forward
```

PostgreSQL будет доступен на `127.0.0.1:5432`.

Закрыть port forwarding:

```sh
make env-port-close
```

## Запуск приложения без Docker

```sh
make todoapp-run
```

Команда:

- выставляет `LOGGER_FOLDER`;
- выставляет `POSTGRES_HOST=localhost`;
- выполняет `go mod tidy`;
- запускает `cmd/todoapp/main.go`.

## Запуск приложения в Docker

```sh
make todoapp-deploy
```

Остановка:

```sh
make todoapp-undeploy
```

## Swagger

Сгенерировать Swagger-артефакты:

```sh
make swagger-gen
```

Swagger UI доступен по адресу:

```text
http://127.0.0.1:5050/swagger/
```

## Очистка

Очистить окружение и данные PostgreSQL:

```sh
make env-cleanup
```

Очистить логи:

```sh
make logs-cleanup
```

Обе команды запрашивают подтверждение перед удалением данных.
