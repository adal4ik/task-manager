package service

import (
	"context"
	"task-manager/internal/core/interfaces/driven"
	"task-manager/internal/utils"

	"github.com/google/uuid"
)

type RegisterService struct {
	repo driven.RegisterDriverInterface
}

func NewRegisterService(repo driven.RegisterDriverInterface) *RegisterService {
	return &RegisterService{
		repo: repo,
	}
}

func (r *RegisterService) RegisterUser(ctx context.Context, login, password, email string) error {
	userID := uuid.New().String()
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	return r.repo.RegisterUser(ctx, login, hashPassword, email, userID)
}
