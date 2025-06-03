package driver

import "context"

type LoginDriverInterface interface {
	LoginUser(ctx context.Context, login, password string) (string, error)
}
