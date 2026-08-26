# Правила участия

Этот документ фиксирует текущий стиль проекта, чтобы новые изменения не
размывали архитектуру.

## Структура фичи

Новая фича должна повторять существующее разбиение:

```text
internal/features/<feature>/transport/http
internal/features/<feature>/service
internal/features/<feature>/repository/postgres
```

Если фича не использует PostgreSQL, имя реализации repository должно отражать
источник данных, например `file_system`.

## Dependency direction

Соблюдайте направление зависимостей:

```text
transport -> service -> repository
```

Transport не должен импортировать PostgreSQL implementation. Service не должен
знать про HTTP DTO. Repository не должен знать про HTTP request/response.

## Интерфейсы

Интерфейсы объявляются рядом с потребителем:

- handler объявляет нужный ему service interface;
- service объявляет нужный ему repository interface.

Не добавляйте общие интерфейсы заранее. Они нужны только когда появляется
реальный потребитель.

## DTO и domain

HTTP DTO должны жить в `transport/http`.

Domain types и domain validation должны жить в `internal/core/domain`, если
эти правила являются бизнес-инвариантами, а не особенностью HTTP.

Mapping между DTO, model и domain держите рядом с соответствующим слоем.

## Ошибки

Для ошибок, влияющих на HTTP-статус, используйте sentinel errors из
`internal/core/errors` и оборачивайте их через `%w`.

Новый тип ошибки стоит добавлять только если response handler должен отличать
ее от существующих случаев.

## Swagger

Новые HTTP endpoints должны иметь Swagger-аннотации рядом с handler method.

После изменения API нужно выполнить:

```sh
make swagger-gen
```

Сгенерированные файлы `docs/swagger.json`, `docs/swagger.yaml` и `docs/docs.go`
должны соответствовать коду.

## Миграции

Новые изменения схемы БД добавляются через миграции:

```sh
make migrate-create seq=descriptive_name
```

Миграции должны иметь both up and down scripts.

## Проверки перед merge

Минимальный checklist:

- код форматирован `gofmt`;
- приложение запускается локально или в Docker;
- миграции применяются на чистую БД;
- Swagger пересобран после изменения API;
- добавлены или обновлены тесты для измененной логики;
- документация обновлена, если изменился публичный контракт.
