package handlers

import (
	"errors"
	"net/http"

	"github.com/abhirup7477/go-inventory-management/internal/delivery/http/dto"
	customerrors "github.com/abhirup7477/go-inventory-management/internal/domain/customerrors"
	"github.com/abhirup7477/go-inventory-management/internal/usecase"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	uc *usecase.CategoriesUsecase
}

func NewCategoryHandler(uc *usecase.CategoriesUsecase) *CategoryHandler {
	return &CategoryHandler{uc: uc}
}

func (h *CategoryHandler) GetCaterogiesHandlerFunc(c *gin.Context) {
	ctx := c.Request.Context()
	categories, err := h.uc.GetCategories(ctx)
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
		return
	}

	res := dto.ToCategoriesResponseList(categories)
	c.JSON(http.StatusOK, res)
}
