package repository

import "database/sql"

type Repository struct {
	AuthRepository *AuthRepository
	TaskRepository *TaskRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		AuthRepository: NewAuthRepository(db),
		TaskRepository: NewTaskRepository(db),
	}
}
