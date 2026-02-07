package handlers

import (
	"errors"
	"net/http"

	"github.com/abhirup7477/go-inventory-management/internal/delivery/http/dto"
	customerrors "github.com/abhirup7477/go-inventory-management/internal/domain/customErrors"
	"github.com/abhirup7477/go-inventory-management/internal/domain/models"
	"github.com/abhirup7477/go-inventory-management/internal/usecase"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	uc *usecase.ProductUsecase
}

func NewProductHandler(uc *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{uc: uc}
}

func (p *ProductHandler) GetAllProducts(c *gin.Context) {
	var products []models.Products
	products, err := p.uc.GetProducts()
	if err != nil {
		if errors.Is(err, customerrors.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
		}
	}

	res := dto.ToProductResponseList(products)
	c.JSON(http.StatusOK, res)
}
