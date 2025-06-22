package repository

import (
	"context"
	"database/sql"
	"task-manager/internal/core/domain/dao"
	"task-manager/internal/utils"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) GetUserByLogin(ctx context.Context, login string) (dao.Users, error) {
	var user dao.Users
	query := "SELECT user_id, login, password FROM users WHERE login = $1"
	err := a.db.QueryRowContext(ctx, query, login).Scan(&user.UserID, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return dao.Users{}, utils.ErrNoRows // User not found
		}
		return dao.Users{}, err // Other error
	}
	return user, nil // User found
}

func (a *AuthRepository) RegisterUser(ctx context.Context, login, hashPassword, email, userID string) error {
	query := `INSERT INTO users (user_id, login, password, email) VALUES ($1, $2, $3, $4)`
	_, err := a.db.ExecContext(ctx, query, userID, login, hashPassword, email)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthRepository) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	query := `SELECT email FROM users WHERE email = $1`
	var existingEmail string
	err := a.db.QueryRowContext(ctx, query, email).Scan(&existingEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err // Return error if any other error occurs
	}
	if existingEmail == "" {
		return false, nil // Email does not exist
	}
	return true, nil // Email exists
}

func (a *AuthRepository) CheckLoginExists(ctx context.Context, login string) (bool, error) {
	query := `SELECT login FROM users WHERE login = $1`
	var existingLogin string
	err := a.db.QueryRowContext(ctx, query, login).Scan(&existingLogin)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err // Return error if any other error occurs
	}
	if existingLogin == "" {
		return false, nil // Login does not exist
	}
	return true, nil // Login exists
}
