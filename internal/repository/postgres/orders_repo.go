package postgres

import (
	"database/sql"
	"fmt"

	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
)

type OrdersRepository struct {
	db *sql.DB
}

func NewOrdersRepository(db *sql.DB) OrdersRepository {
	return OrdersRepository{db: db}
}

func (o OrdersRepository) GetAllOrders() ([]models.Orders, error) {
	var orders []models.Orders

	query := `
		select
			order_id,
			prod_id,
			quantity,
			order_date
		from orders
	`
	rows, err := o.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Query orders failed: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order models.Orders
		rows.Scan(
			&order.OrderId,
			&order.ProdId,
			&order.Quantity,
			&order.OrderDate,
		)
		orders = append(orders, order)
	}

	return orders, nil
}
