package routes

import (
	"database/sql"

	"github.com/abhirup7477/go-inventory-management/internal/delivery/http/handlers"
	"github.com/abhirup7477/go-inventory-management/internal/repository/postgres"
	"github.com/abhirup7477/go-inventory-management/internal/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterTasksRoutes(router *gin.Engine, db *sql.DB) {
	c := postgres.NewCategoriesRepository(db)
	p := postgres.NewProductRepository(db)
	o := postgres.NewOrdersRepository(db)

	uc := usecase.NewTasksUsecase(c, p, o)
	h := handlers.NewTasksHandler(uc)

	router.Group("/tasks")
	{
		router.GET("/inventory", h.GetTasks)
	}
}
