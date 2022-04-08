package repositories

import (
	"database/sql"
)

type PetRepository struct {
	db *sql.DB
}

type (
	Peter interface {
	}
)

func InitPetRepository(db *sql.DB) Peter {
	return &PetRepository{
		db: db,
	}
}
