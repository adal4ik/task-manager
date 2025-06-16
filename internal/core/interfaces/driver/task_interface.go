package driver

import (
	"context"
	"task-manager/internal/core/domain/dto"
)

type TasksDriverInterface interface {
	CreateTask(ctx context.Context, task dto.Task) error
	GetTasks(ctx context.Context, userID string, status string, priority string) ([]dto.Task, error)
	UpdateTask(ctx context.Context, task dto.Task, taskID string) error
	DeleteTask(ctx context.Context, userID string, taskID string) error
}
