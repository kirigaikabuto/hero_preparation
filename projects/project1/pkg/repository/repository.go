package repository

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

func NewRepository() *Repository {
	return &Repository{}
}
