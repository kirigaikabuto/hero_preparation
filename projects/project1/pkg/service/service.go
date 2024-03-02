package service

import "github.com/kirigaikabuto/hero_preparation/projects/project1/pkg/repository"

type Authorization interface {
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
	return &Service{}
}
