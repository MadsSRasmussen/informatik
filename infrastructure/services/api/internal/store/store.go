package store

import (
	"database/sql"
	"errors"
)

type Store struct {
	db *sql.DB
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

var ErrNotFound = errors.New("resource not found")
