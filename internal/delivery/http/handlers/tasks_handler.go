package handlers

import (
	"errors"
	"net/http"

	"github.com/abhirup7477/go-inventory-management/internal/delivery/http/dto"
	"github.com/abhirup7477/go-inventory-management/internal/usecase"
	"github.com/gin-gonic/gin"
)

type TasksHandler struct {
	uc *usecase.TasksUsecase
}

func NewTasksHandler(uc *usecase.TasksUsecase) *TasksHandler {
	return &TasksHandler{uc: uc}
}

func (h *TasksHandler) GetTasks(c *gin.Context) {
	categories, cerr := h.uc.GetCategories()
	products, perr := h.uc.GetProducts()
	orders, oerr := h.uc.GetOrders()

	isCategoriesFound := errors.Is(nil, cerr)
	isProductsFound := errors.Is(nil, perr)
	isOrdersFound := errors.Is(nil, oerr)

	if !isCategoriesFound && !isProductsFound && !isOrdersFound {
		c.JSON(http.StatusNotFound, gin.H{
			"status":    "Not found",
			"error-msg": "No records founds",
		})
		return
	}

	res := make(map[string]interface{})

	res["categories"] = dto.ToCategoriesResponseList(categories)
	res["products"] = dto.ToProductResponseList(products)
	res["orders"] = dto.ToOrdersResponseList(orders)

	c.JSON(http.StatusOK, res)
}
