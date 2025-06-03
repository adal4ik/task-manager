package repository

import (
	"context"
	"database/sql"
	"task-manager/internal/core/domain/dao"
	"task-manager/internal/utils"
)

type LoginRepository struct {
	db *sql.DB
}

func NewLoginRepository(db *sql.DB) *LoginRepository {
	return &LoginRepository{
		db: db,
	}
}

func (l *LoginRepository) GetUserByLogin(ctx context.Context, login string) (dao.Users, error) {
	var user dao.Users
	query := "SELECT user_id, login, password FROM users WHERE login = $1"
	err := l.db.QueryRowContext(ctx, query, login).Scan(&user.UserID, &user.Login, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return dao.Users{}, utils.ErrNoRows // User not found
		}
		return dao.Users{}, err // Other error
	}
	return user, nil // User found
}
