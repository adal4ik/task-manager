package repository

import "database/sql"

type Repository struct {
	RegisterRepository *RegisterRepository
	LoginRepository    *LoginRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		RegisterRepository: NewRegisterRepository(db),
		LoginRepository:    NewLoginRepository(db),
	}
}
