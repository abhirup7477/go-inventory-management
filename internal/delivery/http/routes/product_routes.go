package routes

import (
	"database/sql"

	"github.com/abhirup7477/go-inventory-management/internal/delivery/http/handlers"
	"github.com/abhirup7477/go-inventory-management/internal/repository/postgres"
	"github.com/abhirup7477/go-inventory-management/internal/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterProductsRoutes(router *gin.Engine, db *sql.DB) {
	p := postgres.NewProductRepository(db)
	c := postgres.NewCategoriesRepository(db)
	puc := usecase.NewProductUsecase(p, c)
	h := handlers.NewProductHandler(puc)

	router.Group("/products")
	{
		router.GET("/allProducts", h.GetAllProducts)
	}
}
