package usecase

import (
	customerrors "github.com/abhirup7477/go-inventory-management/internal/domain/customerrors"
	"github.com/abhirup7477/go-inventory-management/internal/domain/interfaces"
	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
)

type ProductUsecase struct {
	productRepo  interfaces.ProductsRepo
	categoryRepo interfaces.CategoriesRepo
}

func NewProductUsecase(p interfaces.ProductsRepo, c interfaces.CategoriesRepo) *ProductUsecase {
	return &ProductUsecase{
		productRepo:  p,
		categoryRepo: c,
	}
}

func (p *ProductUsecase) GetProducts() ([]models.Products, error) {
	products, err := p.productRepo.GetAllProducts()
	if len(products) == 0 {
		return nil, customerrors.ErrNotFound
	}
	return products, err
}
