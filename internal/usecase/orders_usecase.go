package usecase

import (
	"context"

	customerrors "github.com/abhirup7477/go-inventory-management/internal/domain/customerrors"
	"github.com/abhirup7477/go-inventory-management/internal/domain/interfaces"
	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
)

type OrdersUsecase struct {
	orderRepo   interfaces.OrdersRepo
	productRepo interfaces.ProductsRepo
}

func NewOrdersUsecase(o interfaces.OrdersRepo, p interfaces.ProductsRepo) *OrdersUsecase {
	return &OrdersUsecase{
		orderRepo:   o,
		productRepo: p,
	}
}

func (uc *OrdersUsecase) GetOrders(ctx context.Context) ([]models.Orders, error) {
	orders, err := uc.orderRepo.GetAllOrders()
	if len(orders) == 0 {
		return nil, customerrors.ErrNotFound
	}
	return orders, err
}
