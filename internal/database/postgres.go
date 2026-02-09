package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Connection() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		return nil, errors.New("Error loading .env file")
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
		return nil, fmt.Errorf("Failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Failed to ping database: %w", err)
	}

	return db, nil
}
