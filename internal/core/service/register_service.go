package service

import "task-manager/internal/core/interfaces/driven"

type RegisterService struct {
	repo driven.RegisterDriverInterface
}

func NewRegisterService(repo driven.RegisterDriverInterface) *RegisterService {
	return &RegisterService{
		repo: repo,
	}
}
