package driver

import (
	"context"
	"task-manager/internal/core/domain/dto"
)

type TasksDriverInterface interface {
	CreateTask(ctx context.Context, task dto.Task) error
	GetTasks(ctx context.Context, userID string) ([]dto.Task, error)
}
