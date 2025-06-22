package service

import (
	"context"
	"os"
	"task-manager/internal/core/interfaces/driven"
	"task-manager/internal/utils"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo driven.AuthDrivenInterface
}

func NewAuthService(repo driven.AuthDrivenInterface) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) LoginUser(ctx context.Context, login, password string) (string, error) {
	user, err := a.repo.GetUserByLogin(ctx, login)
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

func (a *AuthService) RegisterUser(ctx context.Context, login, password, email string) error {
	userID := uuid.New().String()
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	return a.repo.RegisterUser(ctx, login, hashPassword, email, userID)
}

func (a *AuthService) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	exists, err := a.repo.CheckEmailExists(ctx, email)
	if err != nil {
		return false, err
	}
	return exists, nil
}
func (a *AuthService) CheckLoginExists(ctx context.Context, login string) (bool, error) {
	exists, err := a.repo.CheckLoginExists(ctx, login)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (a *AuthService) LogoutUser(ctx context.Context, userID string) error {

	return nil
}
