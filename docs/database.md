# База данных

Проект использует PostgreSQL. SQL-миграции лежат в `migrations`.

## Schema

Миграция `000001_init.up.sql` создает schema:

```sql
todoapp
```

## Таблица users

```text
todoapp.users
```

| Колонка | Тип | Ограничения |
| --- | --- | --- |
| `id` | `SERIAL` | Primary key. |
| `version` | `BIGINT` | `NOT NULL`, default `1`. |
| `full_name` | `VARCHAR(100)` | `NOT NULL`, длина от 3 до 100. |
| `phone_number` | `VARCHAR(15)` | `NULL`, формат `+` и цифры. |

`phone_number` должен соответствовать выражению:

```text
^\+[0-9]{9,14}$
```

## Таблица tasks

```text
todoapp.tasks
```

| Колонка | Тип | Ограничения |
| --- | --- | --- |
| `id` | `SERIAL` | Primary key. |
| `version` | `BIGINT` | `NOT NULL`, default `1`. |
| `title` | `VARCHAR(100)` | `NOT NULL`, длина от 1 до 100. |
| `description` | `VARCHAR(1000)` | `NULL`, если задано - длина от 1 до 1000. |
| `completed` | `BOOLEAN` | `NOT NULL`. |
| `created_at` | `TIMESTAMPTZ` | `NOT NULL`. |
| `completed_at` | `TIMESTAMPTZ` | `NULL`. |
| `author_user_id` | `INTEGER` | `NOT NULL`, foreign key на `todoapp.users(id)`. |

Инвариант завершения:

- если `completed = false`, `completed_at` должен быть `NULL`;
- если `completed = true`, `completed_at` должен быть задан;
- `completed_at` не может быть раньше `created_at`.

## Version

Поле `version` используется для оптимистичной блокировки в PATCH-операциях.
Репозиторий сначала читает текущую запись, сервис применяет patch к доменной
модели, затем repository делает `UPDATE ... WHERE id = $id AND version = $version`.

Если запись не обновилась, операция считается конфликтом конкурентного доступа.

## Миграции

Миграции запускаются через Docker image `migrate/migrate`.

```sh
make migrate-up
make migrate-down
```

Новые миграции создаются командой:

```sh
make migrate-create seq=descriptive_name
```
