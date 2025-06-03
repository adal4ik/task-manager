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
		return "", err // Handle error from repository
	}
	if user.Login == "" {
		return "", utils.ErrNoRows // Handle case where user is not found
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", utils.ErrInvalidCredentials // Handle invalid password
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // Token expiration time
	})
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		return "", utils.ErrMissingSecret // Handle missing secret key
	}
	tokenString, err := token.SignedString([]byte(secret)) // Replace with your secret key
	if err != nil {
		return "", err // Handle error during token signing
	}
	return tokenString, nil // Return the generated JWT token
}
