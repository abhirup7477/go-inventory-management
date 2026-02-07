package models

import (
	"time"

	"github.com/google/uuid"
)

type Orders struct {
	OrderId   uuid.UUID
	ProdId    uuid.UUID
	Quantity  int
	OrderDate time.Time
}
