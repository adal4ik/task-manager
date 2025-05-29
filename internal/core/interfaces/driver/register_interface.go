package driver

import "context"

type RegisterDriverInterface interface {
	RegisterUser(ctx context.Context, login, hashPassword, email string) error
}
