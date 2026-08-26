# Конфигурация

Конфигурация читается из переменных окружения через `envconfig`. Для локальной
разработки используется `.env`, пример значений находится в `.env.example`.

## Core

| Переменная | Обязательна | По умолчанию | Описание |
| --- | --- | --- | --- |
| `TIME_ZONE` | нет | `UTC` | Timezone приложения. Значение передается в `time.LoadLocation`. |

## HTTP

Префикс `HTTP_` обрабатывается конфигом HTTP-сервера.

| Переменная | Обязательна | По умолчанию | Описание |
| --- | --- | --- | --- |
| `HTTP_ADDR` | да | нет | Адрес HTTP-сервера, например `:5050`. |
| `HTTP_SHUTDOWN_TIMEOUT` | нет | `30s` | Таймаут graceful shutdown. |
| `HTTP_ALLOWED_ORIGINS` | да | нет | Список origin для CORS. |

Пример:

```env
HTTP_ADDR=:5050
HTTP_SHUTDOWN_TIMEOUT=30s
HTTP_ALLOWED_ORIGINS=http://localhost:5050,null
```

## PostgreSQL

Префикс `POSTGRES_` используется реализацией `pgx` pool.

| Переменная | Обязательна | По умолчанию | Описание |
| --- | --- | --- | --- |
| `POSTGRES_HOST` | да | нет | Host PostgreSQL. |
| `POSTGRES_PORT` | нет | `5432` | Port PostgreSQL. |
| `POSTGRES_USER` | да | нет | Имя пользователя. |
| `POSTGRES_PASSWORD` | да | нет | Пароль. |
| `POSTGRES_DB` | да | нет | Имя базы данных. |
| `POSTGRES_TIMEOUT` | да | нет | Таймаут операций репозитория. |

## Logger

Префикс `LOGGER_` используется конфигом логгера.

| Переменная | Обязательна | По умолчанию | Описание |
| --- | --- | --- | --- |
| `LOGGER_LEVEL` | нет | `DEBUG` | Уровень логирования. |
| `LOGGER_FOLDER` | да | нет | Каталог для файлов логов. |

При локальном запуске через `make todoapp-run` `LOGGER_FOLDER` переопределяется
на `out/logs`, а `POSTGRES_HOST` на `localhost`.
