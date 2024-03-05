package service

import (
	"github.com/kirigaikabuto/hero_preparation/projects/project1/models"
	"github.com/kirigaikabuto/hero_preparation/projects/project1/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
}

type Catalog interface {
}

type BookCatalog interface {
}

type Service struct {
	Authorization
	Catalog
	BookCatalog
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
