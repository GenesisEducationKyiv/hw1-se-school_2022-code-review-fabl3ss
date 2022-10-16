package routes

import (
	"orders/src/pkg/delivery/http/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, h *handlers.Handlers) {
	router.POST("/register-customer", h.OrdersHandler.CreateCustomer)
}
