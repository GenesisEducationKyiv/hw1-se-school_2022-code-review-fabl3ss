package routes

import (
	"customers/src/pkg/delivery/http/handlers"

	"github.com/dtm-labs/dtm/dtmutil"
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, h *handlers.Handlers) {
	router.POST(
		"/create-customer",
		dtmutil.WrapHandler2(h.CustomerHandler.CreateCustomer),
	)
	router.POST(
		"/create-customer-compensate",
		dtmutil.WrapHandler2(h.CustomerHandler.CreateCustomerCompensate),
	)
}
