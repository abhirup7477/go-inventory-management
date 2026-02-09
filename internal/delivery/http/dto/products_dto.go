package dto

import (
	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
	"github.com/google/uuid"
)

type CreateProductsRequest struct {
	ProductName string    `json:"product_name" binding:"required"`
	CategoryId  uuid.UUID `json:"category_id" binding:"required"`
	Cost        float64   `json:"cost" binding:"required"`
	Quantity    int       `json:"quantity" binding:"required,gte=0"`
	Description string    `json:"description"`
}

type ProductsResponse struct {
	Id          uuid.UUID `json:"id"`
	ProductName string    `json:"product_name"`
	CategoryId  uuid.UUID `json:"category_id"`
	Cost        float64   `json:"cost"`
	Quantity    int       `json:"quantity"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
}

func ToProductResPonse(p models.Products) ProductsResponse {
	return ProductsResponse{
		Id:          p.Id,
		ProductName: p.ProductName,
		CategoryId:  p.CategoryId,
		Cost:        p.Cost,
		Quantity:    p.Quantity,
		Status:      p.Status,
		Description: p.Description,
	}
}

func ToProductResponseList(products []models.Products) []ProductsResponse {
	var res = make([]ProductsResponse, 0, len(products))
	for _, product := range products {
		res = append(res, ToProductResPonse(product))
	}
	return res
}
