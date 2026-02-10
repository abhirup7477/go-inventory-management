package usecase

import (
	"context"
	"errors"
	"log"

	customerrors "github.com/abhirup7477/go-inventory-management/internal/domain/customerrors"
	"github.com/abhirup7477/go-inventory-management/internal/domain/interfaces"
	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
)

var ErrNotFound = errors.New("Records Not Found")

type TasksUsecase struct {
	categoryRepo interfaces.CategoriesRepo
	productRepo  interfaces.ProductsRepo
	orderRepo    interfaces.OrdersRepo
	mailer       interfaces.Mailer
}

type TasksBundle struct {
	Categories []models.Categories
	Products   []models.Products
	Orders     []models.Orders

	CategoryErr error
	ProductErr  error
	OrderErr    error
}

func NewTasksUsecase(
	c interfaces.CategoriesRepo,
	p interfaces.ProductsRepo,
	o interfaces.OrdersRepo,
	m interfaces.Mailer,
) *TasksUsecase {
	return &TasksUsecase{
		categoryRepo: c,
		productRepo:  p,
		orderRepo:    o,
		mailer:       m,
	}
}

func (uc *TasksUsecase) GetAllTasks(ctx context.Context, email string) (
	*TasksBundle, error,
) {
	categories, cerr := uc.categoryRepo.GetAllCategories()
	products, perr := uc.productRepo.GetAllProducts()
	orders, oerr := uc.orderRepo.GetAllOrders()

	// log.Printf("%+v\n", categories)
	// log.Printf("%+v\n", products)
	// log.Printf("%+v\n", orders)

	isCategoriesNotFound := cerr != nil
	isProductsNotFound := perr != nil
	isOrdersNotFound := oerr != nil

	if isCategoriesNotFound && isProductsNotFound && isOrdersNotFound {
		return nil, ErrNotFound
	}

	go func(email string) {
		bg := context.Background()

		if email == "" {
			log.Println("No email id found!")
		} else {
			err := uc.mailer.SendTasksFetchedEmail(bg, email)
			if err != nil {
				log.Printf("Email failed: %v\n", err)
			}
		}
	}(email)

	return &TasksBundle{
		Categories: categories,
		Products:   products,
		Orders:     orders,

		CategoryErr: cerr,
		ProductErr:  perr,
		OrderErr:    oerr,
	}, nil
}

func (uc *TasksUsecase) GetCategories() ([]models.Categories, error) {
	categories, err := uc.categoryRepo.GetAllCategories()
	if len(categories) == 0 {
		return nil, customerrors.ErrNotFound
	}
	return categories, err
}

func (p *TasksUsecase) GetProducts(ctx context.Context) ([]models.Products, error) {
	products, err := p.productRepo.GetAllProducts()
	if len(products) == 0 {
		return nil, customerrors.ErrNotFound
	}
	return products, err
}

func (uc *TasksUsecase) GetOrders(ctx context.Context) ([]models.Orders, error) {
	orders, err := uc.orderRepo.GetAllOrders()
	if len(orders) == 0 {
		return nil, customerrors.ErrNotFound
	}
	return orders, err
}
