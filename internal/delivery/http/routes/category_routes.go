package routes

import (
	"database/sql"

	"github.com/abhirup7477/go-inventory-management/internal/delivery/http/handlers"
	"github.com/abhirup7477/go-inventory-management/internal/repository/postgres"
	"github.com/abhirup7477/go-inventory-management/internal/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(router *gin.Engine, db *sql.DB) {
	c := postgres.NewCategoriesRepository(db)
	uc := usecase.NewCategoriesUsecase(c)
	h := handlers.NewCategoryHandler(uc)

	router.Group("/categories")
	{
		router.GET("/allCategories", h.GetCaterogiesHandlerFunc)
	}
}
