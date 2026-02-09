package dto

import (
	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
	"github.com/google/uuid"
)

type CreateCategoriesRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoriesResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func ToCategoriesResponse(c models.Categories) CategoriesResponse {
	return CategoriesResponse{
		Id:   c.Id,
		Name: c.Name,
	}
}

func ToCategoriesResponseList(categories []models.Categories) []CategoriesResponse {
	var res []CategoriesResponse
	for _, category := range categories {
		res = append(res, ToCategoriesResponse(category))
	}
	return res
}
