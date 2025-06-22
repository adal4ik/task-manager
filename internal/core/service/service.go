package service

import (
	"task-manager/internal/adapters/driven/database/repository"
)

type Service struct {
	AuthService *AuthService
	TaskService *TaskService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		AuthService: NewAuthService(repo.AuthRepository),
		TaskService: NewTaskService(repo.TaskRepository),
	}
}
