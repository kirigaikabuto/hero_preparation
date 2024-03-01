package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/login", h.logIn)
	}
	api := router.Group("/api")
	{
		catalog := api.Group("/catalog")
		{
			catalog.POST("/", h.createCatalog)
			catalog.GET("/", h.listCatalog)
			catalog.GET("/:id", h.getCatalogById)
			catalog.PUT("/:id", h.updateCatalog)
			catalog.DELETE("/:id", h.deleteCatalog)
			bookCatalog := catalog.Group("/:id/books")
			{
				bookCatalog.POST("/", h.createBookCatalog)
				bookCatalog.GET("/", h.listBookCatalog)
				bookCatalog.GET("/:book_id", h.getBookCatalogById)
				bookCatalog.PUT("/:book_id", h.updateBookCatalog)
				bookCatalog.DELETE("/:book_id", h.deleteBookCatalog)
			}
		}
	}
	return router
}
