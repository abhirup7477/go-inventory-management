package dto

import (
	"time"

	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
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

func ToOrdersResponse(order models.Orders) OrdersResponse {
	return OrdersResponse{
		OrderId:   order.OrderId,
		ProdId:    order.ProdId,
		Quantity:  order.Quantity,
		OrderDate: order.OrderDate,
	}
}

func ToOrdersResponseList(orders []models.Orders) []OrdersResponse {
	var res []OrdersResponse
	for _, order := range orders {
		res = append(res, ToOrdersResponse(order))
	}
	return res
}
