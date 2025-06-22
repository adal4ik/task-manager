package driven

import (
	"context"
	"task-manager/internal/core/domain/dao"
)

type AuthDrivenInterface interface {
	GetUserByLogin(ctx context.Context, login string) (dao.Users, error)
	RegisterUser(ctx context.Context, login, hashPassword, email, userID string) error
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	CheckLoginExists(ctx context.Context, login string) (bool, error)
}
