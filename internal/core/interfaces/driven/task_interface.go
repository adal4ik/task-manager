package driven

import (
	"context"
	"task-manager/internal/core/domain/dto"
)

type TasksDrivenInteface interface {
	CreateTask(ctx context.Context, task dto.Task) error
}
