package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/models"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user models.User) (int, error) {
	return 0, nil
}
