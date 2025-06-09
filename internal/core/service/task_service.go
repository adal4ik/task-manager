package service

import (
	"context"
	"task-manager/internal/core/domain/dto"
	"task-manager/internal/core/interfaces/driven"
)

type TaskService struct {
	repo driven.TasksDrivenInteface
}

func NewTaskService(repo driven.TasksDrivenInteface) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (t *TaskService) CreateTask(ctx context.Context, task dto.Task) error {
	return t.repo.CreateTask(ctx, task)
}
