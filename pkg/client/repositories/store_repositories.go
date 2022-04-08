package repositories

import (
	"database/sql"
)

type StoreRepository struct {
	db *sql.DB
}

type (
	Storeer interface {
	}
)

func InitStoreRepository(db *sql.DB) Storeer {
	return &StoreRepository{
		db: db,
	}
}
