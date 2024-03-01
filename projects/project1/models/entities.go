package models

type Catalog struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UserCatalog struct {
	Id        int `json:"id"`
	UserId    int `json:"user_id"`
	CatalogId int `json:"catalog_id"`
}

type BookCatalog struct {
	Id        int `json:"id"`
	CatalogId int `json:"catalog_id"`
	BookId    int `json:"book_id"`
}

type Book struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
