package driven

import (
	"context"
	"task-manager/internal/core/domain/dao"
)

type TasksDrivenInterface interface {
	CreateTask(ctx context.Context, taskDao dao.Tasks) error
	GetTasks(ctx context.Context, userID string) ([]dao.Tasks, error)
	UpdateTask(ctx context.Context, taskDao dao.Tasks, taskID string) error
}
