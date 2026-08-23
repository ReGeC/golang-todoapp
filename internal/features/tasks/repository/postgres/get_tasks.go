package tasks_postgres_repository

import (
	"context"
	"fmt"

	"github.com/ReGeC/golang-todoapp/internal/core/domain"
)

func (r *TasksRepository) GetTasks(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OpTimeout())
	defer cancel()

	args := []any{limit, offset}

	query := `
	SELECT
		id,
		version,
		title,
		description,
		completed,
		created_at,
		completed_at,
		author_user_id
	FROM todoapp.tasks
	`
	if userID != nil {
		query += `
		WHERE author_user_id = $3
		`

		args = append(args, *userID)
	}
	query += `
	ORDER BY id ASC
	LIMIT $1
	OFFSET $2;
	`

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("select tasks: %w", err)
	}
	defer rows.Close()

	var taskModels []TaskModel
	for rows.Next() {
		var taskModel TaskModel
		err := rows.Scan(
			&taskModel.ID,
			&taskModel.Version,
			&taskModel.Title,
			&taskModel.Description,
			&taskModel.Completed,
			&taskModel.CreatedAt,
			&taskModel.CompletedAt,
			&taskModel.AuthorUserID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan error: %w", err)
		}

		taskModels = append(taskModels, taskModel)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	taskDomains := arrayTaskDomainsFromModels(taskModels)

	return taskDomains, nil
}