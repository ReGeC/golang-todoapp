package tasks_transport_http

import (
	"time"

	"github.com/ReGeC/golang-todoapp/internal/core/domain"
)

type TaskDTOResponse struct {
	ID           int        `json:"id"             example:"5"`
	Version      int        `json:"version"        example:"3"`
	Title        string     `json:"title"          example:"test_task"`
	Description  *string    `json:"description"    example:"test_desc"`
	Completed    bool       `json:"completed"      example:"true"`
	CreatedAt    time.Time  `json:"created_at"     example:"2006-01-26T10:30:20.299941Z"`
	CompletedAt  *time.Time `json:"completed_at"   example:"null"`
	AuthorUserID int        `json:"author_user_id" example:"2"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID: task.ID,
		Version: task.Version,
		Title: task.Title,
		Description: task.Description,
		Completed: task.Completed,
		CreatedAt: task.CreatedAt,
		CompletedAt: task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func arrayTaskDTOFromDomains(tasks []domain.Task) []TaskDTOResponse {
	arrayTaskDTO := make([]TaskDTOResponse, len(tasks))

	for i, task := range tasks {
		arrayTaskDTO[i] = taskDTOFromDomain(task)
	}

	return arrayTaskDTO
}