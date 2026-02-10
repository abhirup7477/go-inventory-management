package usecase

import (
	"context"

	customerrors "github.com/abhirup7477/go-inventory-management/internal/domain/customerrors"
	"github.com/abhirup7477/go-inventory-management/internal/domain/interfaces"
	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
	"github.com/abhirup7477/go-inventory-management/internal/repository/postgres"
)

type CategoriesUsecase struct {
	categoryRepo interfaces.CategoriesRepo
}

func NewCategoriesUsecase(c postgres.CategoriesRepository) *CategoriesUsecase {
	return &CategoriesUsecase{categoryRepo: c}
}

func (uc *CategoriesUsecase) GetCategories(ctx context.Context) ([]models.Categories, error) {
	categories, err := uc.categoryRepo.GetAllCategories()
	if len(categories) == 0 {
		return nil, customerrors.ErrNotFound
	}
	return categories, err
}
