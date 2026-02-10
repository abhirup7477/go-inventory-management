package handlers

import (
	"errors"
	"net/http"

	"github.com/abhirup7477/go-inventory-management/internal/delivery/http/dto"
	customerrors "github.com/abhirup7477/go-inventory-management/internal/domain/customerrors"
	"github.com/abhirup7477/go-inventory-management/internal/usecase"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	uc *usecase.OrdersUsecase
}

func NewOrderHandler(uc *usecase.OrdersUsecase) *OrderHandler {
	return &OrderHandler{uc: uc}
}

func (h *OrderHandler) GetAllOrdersHandlerFunc(c *gin.Context) {
	ctx := c.Request.Context()
	orders, err := h.uc.GetOrders(ctx)
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

	res := dto.ToOrdersResponseList(orders)
	c.JSON(http.StatusOK, res)
}
