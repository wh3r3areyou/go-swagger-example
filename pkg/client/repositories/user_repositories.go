package repositories

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

type (
	Userer interface {
	}
)

func InitUserRepository(db *sql.DB) Userer {
	return &UserRepository{
		db: db,
	}
}
