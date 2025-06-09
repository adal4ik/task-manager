package driver

import (
	"context"
	"task-manager/internal/core/domain/dto"
)

type TasksDriverInterface interface {
	CreateTask(ctx context.Context, task dto.Task) error
}
