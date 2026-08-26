# API

Основной REST API доступен под префиксом:

```text
/api/v1
```

Актуальная машинная спецификация генерируется в Swagger:

```text
/swagger/
/swagger/doc.json
```

Этот документ описывает поведение API на уровне продукта. Детали схем DTO,
статусов и параметров нужно сверять со Swagger.

## Users

Ресурс `users` управляет пользователями.

| Метод | Путь | Назначение |
| --- | --- | --- |
| `POST` | `/api/v1/users` | Создать пользователя. |
| `GET` | `/api/v1/users` | Получить список пользователей. |
| `GET` | `/api/v1/users/{id}` | Получить пользователя по id. |
| `PATCH` | `/api/v1/users/{id}` | Частично изменить пользователя. |
| `DELETE` | `/api/v1/users/{id}` | Удалить пользователя. |

Поля пользователя:

- `id`;
- `version`;
- `full_name`;
- `phone_number`.

`phone_number` может быть `null`.

## Tasks

Ресурс `tasks` управляет задачами.

| Метод | Путь | Назначение |
| --- | --- | --- |
| `POST` | `/api/v1/tasks` | Создать задачу. |
| `GET` | `/api/v1/tasks` | Получить список задач. |
| `GET` | `/api/v1/tasks/{id}` | Получить задачу по id. |
| `PATCH` | `/api/v1/tasks/{id}` | Частично изменить задачу. |
| `DELETE` | `/api/v1/tasks/{id}` | Удалить задачу. |

Поля задачи:

- `id`;
- `version`;
- `title`;
- `description`;
- `completed`;
- `created_at`;
- `completed_at`;
- `completion_duration`;
- `author_user_id`.

`description` и `completed_at` могут быть `null`.

## Statistics

Ресурс `statistics` возвращает агрегированную статистику по задачам.

| Метод | Путь | Назначение |
| --- | --- | --- |
| `GET` | `/api/v1/statistics` | Получить статистику по задачам. |

Поддерживаются query-параметры:

- `user_id` - фильтр по автору задачи;
- `from` - начало периода включительно, формат `YYYY-MM-DD`;
- `to` - конец периода не включительно, формат `YYYY-MM-DD`.

Ответ содержит:

- `tasks_created`;
- `tasks_completed`;
- `tasks_completed_rate`;
- `tasks_average_completion_time`.

## PATCH three-state logic

PATCH-запросы используют three-state logic для nullable-полей:

1. Поле отсутствует в JSON: значение не меняется.
2. Поле передано со значением: значение обновляется.
3. Поле передано как `null`: значение очищается, если доменный контракт это разрешает.

Для `users`:

- `full_name` не может быть `null`;
- `phone_number` может быть `null`.

Для `tasks`:

- `title` не может быть `null`;
- `completed` не может быть `null`;
- `description` может быть `null`.

## Ошибки

Общий response handler переводит доменные ошибки в HTTP-статусы:

| Доменная ошибка | HTTP-статус |
| --- | --- |
| `ErrInvalidArguments` | `400 Bad Request` |
| `ErrNotFound` | `404 Not Found` |
| `ErrConflict` | `409 Conflict` |
| прочие ошибки | `500 Internal Server Error` |

Формат ошибки:

```json
{
  "error": "low-level error text",
  "message": "operation context"
}
```
