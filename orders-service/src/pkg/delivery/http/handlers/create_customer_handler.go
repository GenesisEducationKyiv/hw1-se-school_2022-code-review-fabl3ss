package handlers

import (
	"net/http"
	"orders/src/pkg/domain/models"
	"orders/src/pkg/domain/usecases"

	"github.com/gin-gonic/gin"
)

func NewOrdersHandler(orderUsecase usecases.OrdersUsecase) *OrdersHandler {
	return &OrdersHandler{
		ordersUsecase: orderUsecase,
	}
}

type OrdersHandler struct {
	ordersUsecase usecases.OrdersUsecase
}

func (h *OrdersHandler) CreateCustomer(c *gin.Context) {
	customer := new(models.Recipient)
	if err := c.BindJSON(customer); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	gid, err := h.ordersUsecase.RegisterCustomer(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"gid": gid})
	} else {
		c.JSON(http.StatusOK, gin.H{"gid": gid})
	}
}
