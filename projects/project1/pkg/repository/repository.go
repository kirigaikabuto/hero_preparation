package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
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
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
