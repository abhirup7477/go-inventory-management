package routes

import (
	"database/sql"

	"github.com/abhirup7477/go-inventory-management/internal/delivery/http/handlers"
	"github.com/abhirup7477/go-inventory-management/internal/repository/postgres"
	"github.com/abhirup7477/go-inventory-management/internal/usecase"
	"github.com/gin-gonic/gin"
)

func RegisterOrdersRoutes(router *gin.Engine, db *sql.DB) {
	o := postgres.NewOrdersRepository(db)
	p := postgres.NewProductRepository(db)
	uc := usecase.NewOrdersUsecase(o, p)
	h := handlers.NewOrderHandler(uc)

	r := router.Group("/inventory/orders")
	{
		r.GET("/allOrders", h.GetAllOrdersHandlerFunc)
	}
}
