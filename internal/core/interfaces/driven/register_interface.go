package driven

import "context"

type RegisterDrivenInterface interface {
	RegisterUser(ctx context.Context, login, hashPassword, email, userID string) error
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	CheckLoginExists(ctx context.Context, login string) (bool, error)
}
