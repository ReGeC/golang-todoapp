package tasks_service

import (
	"context"
	"fmt"

	"github.com/ReGeC/golang-todoapp/internal/core/domain"
	core_errors "github.com/ReGeC/golang-todoapp/internal/core/errors"
)

func (s *TasksService) GetTasks(
	ctx context.Context,
	userID *int,
	limit *int,
	offset *int,
) ([]domain.Task, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"`limit` must be non-negative: %w",
			core_errors.ErrInvalidArguments,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"`offset` must be non-negative: %w",
			core_errors.ErrInvalidArguments,
		)
	}

	taskDomains, err := s.tasksRepository.GetTasks(ctx, userID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("get tasks from repository: %w", err)
	}

	return taskDomains, nil
}