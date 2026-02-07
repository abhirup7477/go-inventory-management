package interfaces

import (
	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
)

type ProductsRepo interface {
	GetAllProducts() ([]models.Products, error)
}

type CategoriesRepo interface {
	GetAllCategories() ([]models.Categories, error)
}

type OrdersRepo interface {
	GetAllOrders() ([]models.Orders, error)
}
