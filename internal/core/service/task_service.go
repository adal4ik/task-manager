package service

import (
	"context"
	"task-manager/internal/core/domain/dto"
	"task-manager/internal/core/interfaces/driven"
)

type TaskService struct {
	repo driven.TasksDrivenInterface
}

func NewTaskService(repo driven.TasksDrivenInterface) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (t *TaskService) CreateTask(ctx context.Context, task dto.Task) error {
	var taskDao, err = dto.TaskToDao(task)
	if err != nil {
		return err
	}
	return t.repo.CreateTask(ctx, taskDao)
}

func (t *TaskService) GetTasks(ctx context.Context, userID string) ([]dto.Task, error) {
	tasks, err := t.repo.GetTasks(ctx, userID)
	if err != nil {
		return nil, err
	}
	var taskDtos []dto.Task
	for _, task := range tasks {
		taskDtos = append(taskDtos, dto.DaoToTask(task))
	}
	return taskDtos, nil
}
