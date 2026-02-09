package usecase

import (
	customerrors "github.com/abhirup7477/go-inventory-management/internal/domain/customerrors"
	"github.com/abhirup7477/go-inventory-management/internal/domain/interfaces"
	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
)

type TasksUsecase struct {
	categoryRepo interfaces.CategoriesRepo
	productRepo  interfaces.ProductsRepo
	orderRepo    interfaces.OrdersRepo
}

func NewTasksUsecase(
	c interfaces.CategoriesRepo,
	p interfaces.ProductsRepo,
	o interfaces.OrdersRepo,
) *TasksUsecase {
	return &TasksUsecase{
		categoryRepo: c,
		productRepo:  p,
		orderRepo:    o,
	}
}

func (uc *TasksUsecase) GetTasks() (
	[]models.Categories, []models.Products, []models.Orders,
) {
	categories, _ := uc.categoryRepo.GetAllCategories()
	products, _ := uc.productRepo.GetAllProducts()
	orders, _ := uc.orderRepo.GetAllOrders()

	return categories, products, orders
}

func (uc *TasksUsecase) GetCategories() ([]models.Categories, error) {
	categories, err := uc.categoryRepo.GetAllCategories()
	if len(categories) == 0 {
		return nil, customerrors.ErrNotFound
	}
	return categories, err
}

func (p *TasksUsecase) GetProducts() ([]models.Products, error) {
	products, err := p.productRepo.GetAllProducts()
	if len(products) == 0 {
		return nil, customerrors.ErrNotFound
	}
	return products, err
}

func (uc *TasksUsecase) GetOrders() ([]models.Orders, error) {
	orders, err := uc.orderRepo.GetAllOrders()
	if len(orders) == 0 {
		return nil, customerrors.ErrNotFound
	}
	return orders, err
}
