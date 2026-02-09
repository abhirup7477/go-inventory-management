package main

import (
	"fmt"
	"log"

	"github.com/abhirup7477/go-inventory-management/internal/database"
	"github.com/abhirup7477/go-inventory-management/internal/delivery/http/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	db, err := database.Connection()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Connected!")

	router := gin.Default()

	routes.RegisterProductsRoutes(router, db)
	routes.RegisterCategoryRoutes(router, db)
	routes.RegisterOrdersRoutes(router, db)

	routes.RegisterTasksRoutes(router, db)

	router.Run(":8080")
}
