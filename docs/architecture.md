# Архитектура

Проект построен как небольшой монолитный HTTP-сервис с разделением по фичам.
Основная идея: общая инфраструктура находится в `internal/core`, а прикладные
сценарии разнесены по `internal/features`.

## Точка входа

`cmd/todoapp/main.go` выполняет ручную сборку зависимостей:

- читает конфигурацию;
- настраивает timezone;
- создает логгер;
- открывает PostgreSQL pool;
- собирает repository, service и HTTP handler для каждой фичи;
- регистрирует маршруты API, web-страницу и Swagger;
- запускает HTTP-сервер с graceful shutdown.

В проекте нет DI-контейнера. Это осознанно сохраняет запуск прозрачным:
по `main.go` видно весь runtime graph приложения.

## Слои

Типовая фича устроена так:

```text
transport/http -> service -> repository/postgres -> database
```

`transport/http` отвечает за HTTP-контракт: DTO, path/query параметры, JSON,
Swagger-аннотации и преобразование ответа.

`service` содержит бизнес-логику и доменную валидацию. Сервис зависит от
репозитория через интерфейс, объявленный рядом с сервисом.

`repository/postgres` содержит SQL и преобразование database model в domain
model. Репозитории используют общий интерфейс пула из `internal/core`.

`internal/core/domain` содержит доменные сущности и инварианты, которые должны
оставаться валидными независимо от транспорта и базы данных.

## Интерфейсы

Интерфейсы объявляются на стороне потребителя:

- HTTP handler объявляет интерфейс сервиса, который ему нужен.
- Service объявляет интерфейс репозитория, который ему нужен.
- PostgreSQL implementation подключается в `main.go`.

Такой стиль снижает связность и упрощает подмену зависимостей в тестах.

## Общая инфраструктура

`internal/core/transport/http` содержит:

- маршрутизацию и API version prefix;
- middleware chain;
- request parsing helpers;
- response handler;
- error-to-status mapping.

`internal/core/repository/postgres/pool` содержит абстракцию пула. Реализация
на `pgx` лежит в `internal/core/repository/postgres/pool/pgx`.

## Доменные инварианты

Важные правила зафиксированы в domain layer и дополнительно защищены БД:

- `User.FullName` имеет длину от 3 до 100 символов.
- `User.PhoneNumber` либо `NULL`, либо соответствует формату `+` и цифры.
- `Task.Title` имеет длину от 1 до 100 символов.
- `Task.Description` либо `NULL`, либо от 1 до 1000 символов.
- если задача завершена, `CompletedAt` должен быть задан;
- если задача не завершена, `CompletedAt` должен быть `nil`;
- `CompletedAt` не может быть раньше `CreatedAt`.

## Оптимистичная блокировка

В таблицах есть поле `version`. PATCH-операции обновляют запись с условием
`WHERE id = ... AND version = ...`, после чего увеличивают `version`.
Если запись не обновилась, репозиторий возвращает conflict error.

## Расширение новой фичей

Для новой фичи стоит повторять существующую структуру:

```text
internal/features/<feature>/transport/http
internal/features/<feature>/service
internal/features/<feature>/repository/postgres
```

Новые доменные сущности и общие типы добавляются в `internal/core/domain`
только если они действительно разделяются между фичами.
