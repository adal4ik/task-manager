package repository

import "database/sql"

type RegisterRepository struct {
	db *sql.DB
}

func NewRegisterRepository(db *sql.DB) *RegisterRepository {
	return &RegisterRepository{
		db: db,
	}
}
