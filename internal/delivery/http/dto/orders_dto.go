package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateOrdersRequest struct {
	ProdId   uuid.UUID `json:"prod_id"`
	Quantity int       `json:"quantity" binding:"omitempty,gte=1"`
}

type OrdersResponse struct {
	OrderId   uuid.UUID `json:"order_id"`
	ProdId    uuid.UUID `json:"prod_id"`
	Quantity  int       `json:"quantity"`
	OrderDate time.Time `json:"order_date"`
}
