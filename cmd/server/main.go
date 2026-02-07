package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/abhirup7477/go-inventory-management/internal/delivery/http/handlers"
	"github.com/abhirup7477/go-inventory-management/internal/repository/postgres"
	"github.com/abhirup7477/go-inventory-management/internal/usecase"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// var db_url = "postgres://inventory_user:sunu@localhost:5432/inventory_db?sslmode=disable"
	db_url := fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SERVER"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var db *sql.DB
	db, err := sql.Open("pgx", db_url)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("DB connect")

	r := gin.Default()

	p := postgres.NewProductRepository(db)
	c := postgres.NewCategoriesRepository(db)
	puc := usecase.NewProductUsecase(p, c)
	h := handlers.NewProductHandler(puc)

	r.GET("/products", h.GetAllProducts)

	r.Run(":8080")
}
