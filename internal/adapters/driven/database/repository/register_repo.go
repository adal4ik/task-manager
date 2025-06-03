package repository

import (
	"context"
	"database/sql"
)

type RegisterRepository struct {
	db *sql.DB
}

func NewRegisterRepository(db *sql.DB) *RegisterRepository {
	return &RegisterRepository{
		db: db,
	}
}

func (r *RegisterRepository) RegisterUser(ctx context.Context, login, hashPassword, email, userID string) error {
	query := `INSERT INTO users (user_id, login, password, email) VALUES ($1, $2, $3, $4)`
	_, err := r.db.ExecContext(ctx, query, userID, login, hashPassword, email)
	if err != nil {
		return err
	}
	return nil
}

func (r *RegisterRepository) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	query := `SELECT email FROM users WHERE email = $1`
	var existingEmail string
	err := r.db.QueryRowContext(ctx, query, email).Scan(&existingEmail)
	if err != nil {
		return false, err
	}
	if existingEmail == "" {
		return false, nil // Email does not exist
	}
	return true, nil // Email exists
}

func (r *RegisterRepository) CheckLoginExists(ctx context.Context, login string) (bool, error) {
	query := `SELECT login FROM users WHERE login = $1`
	var existingLogin string
	err := r.db.QueryRowContext(ctx, query, login).Scan(&existingLogin)
	if err != nil {
		return false, err
	}
	if existingLogin == "" {
		return false, nil // Login does not exist
	}
	return true, nil // Login exists
}
