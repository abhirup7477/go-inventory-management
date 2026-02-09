package postgres

import (
	"database/sql"
	"fmt"

	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
)

type CategoriesRepository struct {
	db *sql.DB
}

func NewCategoriesRepository(db *sql.DB) CategoriesRepository {
	return CategoriesRepository{db: db}
}

func (c CategoriesRepository) GetAllCategories() ([]models.Categories, error) {
	var categories []models.Categories

	query := `
		select 
			id, 
			name
		from categories
	`
	rows, err := c.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Query categories failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Categories
		rows.Scan(
			&category.Id,
			&category.Name,
		)
		categories = append(categories, category)
	}

	return categories, nil
}
