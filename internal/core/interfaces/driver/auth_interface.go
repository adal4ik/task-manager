package driver

import "context"

type AuthDriverInterface interface {
	LoginUser(ctx context.Context, login, password string) (string, error)
	RegisterUser(ctx context.Context, login, hashPassword, email string) error
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	CheckLoginExists(ctx context.Context, login string) (bool, error)
	LogoutUser(ctx context.Context, userID string) error
}
