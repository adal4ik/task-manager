package driven

import (
	"context"
	"task-manager/internal/core/domain/dao"
)

type TasksDrivenInterface interface {
	CreateTask(ctx context.Context, taskDao dao.Tasks) error
	GetTasks(ctx context.Context, userID string, status string, priority string, search string, orderBy string, sortBy string) ([]dao.Tasks, error)
	UpdateTask(ctx context.Context, taskDao dao.Tasks, taskID string) error
	DeleteTask(ctx context.Context, userID string, taskID string) error
	UpdateTaskStatus(ctx context.Context, userID string, taskID string, status string) error
}
