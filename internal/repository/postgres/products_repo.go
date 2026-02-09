package postgres

import (
	"database/sql"
	"fmt"

	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
)

type ProductsRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductsRepository {
	return &ProductsRepository{db: db}
}

func (p *ProductsRepository) GetAllProducts() ([]models.Products, error) {
	var products []models.Products

	query := `
		select 
			id, 
			product_name, 
			category_id, 
			cost, 
			quantity, 
			status, 
			description
		from products
	`
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Query products: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var product models.Products
		rows.Scan(
			&product.Id,
			&product.ProductName,
			&product.CategoryId,
			&product.Cost,
			&product.Quantity,
			&product.Status,
			&product.Description,
		)
		products = append(products, product)
	}

	return products, nil
}
