package driven

import "context"

type RegisterDriverInterface interface {
	RegisterUser(ctx context.Context, login, hashPassword, email, userID string) error
}
