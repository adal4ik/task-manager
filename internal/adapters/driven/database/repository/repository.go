package repository

import "database/sql"

type Repository struct {
	RegisterRepository *RegisterRepository
	LoginRepository    *LoginRepository
	TaskRepository     *TaskRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		RegisterRepository: NewRegisterRepository(db),
		LoginRepository:    NewLoginRepository(db),
		TaskRepository:     NewTaskRepository(db),
	}
}
