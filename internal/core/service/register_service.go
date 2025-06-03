package service

import (
	"context"
	"task-manager/internal/core/interfaces/driven"
	"task-manager/internal/utils"

	"github.com/google/uuid"
)

type RegisterService struct {
	repo driven.RegisterDrivenInterface
}

func NewRegisterService(repo driven.RegisterDrivenInterface) *RegisterService {
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

func (r *RegisterService) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	exists, err := r.repo.CheckEmailExists(ctx, email)
	if err != nil {
		return false, err
	}
	return exists, nil
}
func (r *RegisterService) CheckLoginExists(ctx context.Context, login string) (bool, error) {
	exists, err := r.repo.CheckLoginExists(ctx, login)
	if err != nil {
		return false, err
	}
	return exists, nil
}
