# Эксплуатация

## Запуск

Приложение может запускаться локально через `go run` или в Docker Compose.

Локальный запуск:

```sh
make todoapp-run
```

Docker-запуск:

```sh
make todoapp-deploy
```

## HTTP

HTTP-сервер использует стандартный `net/http`.

API регистрируется под `/api/v1`. Swagger доступен отдельно:

```text
/swagger/
/swagger/doc.json
```

Также регистрируются маршруты web-фичи для выдачи статической HTML-страницы.

## Middleware

Глобальная цепочка middleware задается в `cmd/todoapp/main.go`:

1. CORS.
2. Request ID.
3. Logger context.
4. Trace logging.
5. Panic recovery.

`RequestID` использует заголовок:

```text
X-REQUEST-ID
```

Если клиент не передал request id, приложение генерирует новый UUID и
возвращает его в response headers.

## Логи

Логгер настраивается через:

- `LOGGER_LEVEL`;
- `LOGGER_FOLDER`.

При Docker-запуске логи пишутся в `/app/out/logs`, который мапится на
локальный каталог `./out/logs`.

## Graceful shutdown

Приложение слушает `SIGINT` и `SIGTERM`. После получения сигнала HTTP-сервер
останавливается через `server.Shutdown` с таймаутом `HTTP_SHUTDOWN_TIMEOUT`.

Если graceful shutdown завершается ошибкой, приложение принудительно закрывает
сервер через `server.Close`.

## PostgreSQL

PostgreSQL запускается сервисом `todoapp-postgres`.

Данные хранятся в:

```text
./out/pgdata
```

Для доступа к базе с хоста можно поднять port forward:

```sh
make env-port-forward
```

## Очистка данных

Очистка базы и логов выполняется отдельными командами:

```sh
make env-cleanup
make logs-cleanup
```

Обе команды требуют ручного подтверждения.
