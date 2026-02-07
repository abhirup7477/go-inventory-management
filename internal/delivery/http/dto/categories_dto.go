package dto

import "github.com/google/uuid"

type CreateCategoriesRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoriesResponse struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
