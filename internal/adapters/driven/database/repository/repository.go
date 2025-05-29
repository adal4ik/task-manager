package repository

import "database/sql"

type Repository struct {
	RegisterRepo *RegisterRepository
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{RegisterRepo: NewRegisterRepository(db)}
}
