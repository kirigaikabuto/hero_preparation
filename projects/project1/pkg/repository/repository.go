package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Catalog interface {
}

type BookCatalog interface {
}

type Repository struct {
	Authorization
	Catalog
	BookCatalog
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
