package usecase

import (
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
	return p.productRepo.GetAllProducts()
}
