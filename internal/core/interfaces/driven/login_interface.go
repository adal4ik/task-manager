package driven

import (
	"context"
	"task-manager/internal/core/domain/dao"
)

type LoginDrivenInterface interface {
	GetUserByLogin(ctx context.Context, login string) (dao.Users, error)
}
