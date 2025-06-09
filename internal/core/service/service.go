package service

import (
	"task-manager/internal/adapters/driven/database/repository"
)

type Service struct {
	RegisterService *RegisterService
	LoginService    *LoginService
	TaskService     *TaskService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		RegisterService: NewRegisterService(repo.RegisterRepository),
		LoginService:    NewLoginService(repo.LoginRepository),
		TaskService:     NewTaskService(repo.TaskRepository),
	}
}
