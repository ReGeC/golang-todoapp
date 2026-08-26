# Тестирование

На момент подготовки документации в проекте нет файлов `*_test.go`.
Ниже описана целевая стратегия тестирования, совместимая с текущей архитектурой.

## Domain unit tests

Покрыть `internal/core/domain`:

- валидацию `User`;
- валидацию `Task`;
- `ApplyPatch` для `UserPatch`;
- `ApplyPatch` для `TaskPatch`;
- `CompletionDuration`;
- поведение nullable-полей.

Эти тесты должны быть быстрыми и не требовать внешних зависимостей.

## Service tests

Сервисы стоит тестировать с fake repository, потому что интерфейсы объявлены на
стороне service layer.

Проверить:

- что сервис валидирует доменную модель до записи;
- что ошибки репозитория оборачиваются без потери sentinel errors;
- что PATCH сначала читает сущность, применяет patch, затем сохраняет;
- что статистика корректно считает rate и average completion time.

## HTTP handler tests

HTTP layer можно тестировать через `httptest`.

Проверить:

- decode и validation ошибок;
- path/query parsing;
- status codes;
- JSON response shape;
- PATCH three-state logic на уровне JSON.

Для handler tests сервисы лучше подменять fake implementation.

## Repository integration tests

Репозитории требуют PostgreSQL. Их стоит запускать как integration tests:

- поднять тестовую БД;
- применить миграции;
- подготовить fixtures;
- выполнить repository methods;
- проверить SQL-ограничения и mapping ошибок.

Особенно важны:

- foreign key violation при создании задачи с несуществующим автором;
- optimistic locking conflict в PATCH;
- `ErrNoRows` для get/delete;
- фильтрация и пагинация.

## Smoke tests

Минимальный smoke test для запущенного приложения:

1. Создать пользователя.
2. Создать задачу.
3. Получить задачу по id.
4. Изменить `completed`.
5. Получить статистику.
6. Удалить задачу и пользователя.

## Рекомендуемый порядок внедрения

1. Domain unit tests.
2. Service tests.
3. HTTP handler tests.
4. Repository integration tests.
5. End-to-end smoke scenario.
