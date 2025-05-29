package service

import (
	"task-manager/internal/adapters/driven/database/repository"
)

type Service struct {
	RegisterService *RegisterService
}

func NewService(repo *repository.Repository) *Service {
	return &Service{RegisterService: NewRegisterService(repo.RegisterRepo)}
}
