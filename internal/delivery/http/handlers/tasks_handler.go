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
	ctx := c.Request.Context()
	email := c.Request.Header.Get("email")
	bundle, err := h.uc.GetAllTasks(ctx, email)

	if errors.Is(err, usecase.ErrNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":    "Not found",
			"error-msg": "No records founds",
		})
		return
	}

	res := gin.H{}

	if bundle.CategoryErr != nil {
		res["categories"] = "Failed to fetch categories"
	} else {
		res["categories"] = dto.ToCategoriesResponseList(bundle.Categories)
	}

	if bundle.ProductErr != nil {
		res["products"] = "Failed to fetch products"
	} else {
		res["products"] = dto.ToProductResponseList(bundle.Products)
	}

	if bundle.OrderErr != nil {
		res["orders"] = "Failed to fetch orders"
	} else {
		res["orders"] = dto.ToOrdersResponseList(bundle.Orders)
	}

	c.JSON(http.StatusOK, res)
}
