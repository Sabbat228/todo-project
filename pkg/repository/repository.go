package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type TodoLists interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoLists
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
