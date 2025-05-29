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
