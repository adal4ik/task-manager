package driver

import "context"

type RegisterDriverInterface interface {
	RegisterUser(ctx context.Context, login, hashPassword, email string) error
	CheckEmailExists(ctx context.Context, email string) (bool, error)
	CheckLoginExists(ctx context.Context, login string) (bool, error)
}
