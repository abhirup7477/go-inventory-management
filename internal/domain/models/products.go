package models

import "github.com/google/uuid"

type Products struct {
	Id          uuid.UUID
	ProductName string
	CategoryId  uuid.UUID
	Cost        float64
	Quantity    int
	Status      string
	Description string
}
