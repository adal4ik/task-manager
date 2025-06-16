package service

import (
	"context"
	"os"
	"task-manager/internal/core/interfaces/driven"
	"task-manager/internal/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	repo driven.LoginDrivenInterface
}

func NewLoginService(repo driven.LoginDrivenInterface) *LoginService {
	return &LoginService{
		repo: repo,
	}
}

func (l *LoginService) LoginUser(ctx context.Context, login, password string) (string, error) {
	user, err := l.repo.GetUserByLogin(ctx, login)
	if err != nil {
		return "", err
	}
	if user.Login == "" {
		return "", utils.ErrNoRows
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", utils.ErrInvalidCredentials
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return "", utils.ErrMissingSecret
	}
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
