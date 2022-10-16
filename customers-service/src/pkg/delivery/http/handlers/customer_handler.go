package handlers

import (
	"customers/src/pkg/domain/models"
	"customers/src/pkg/domain/usecases"

	"github.com/dtm-labs/client/dtmcli"
	"github.com/gin-gonic/gin"
)

const gid = "gid"

func NewCustomerHandler(customer usecases.CustomerUsecase) *CustomerHandler {
	return &CustomerHandler{
		customerUsecase: customer,
	}
}

type CustomerHandler struct {
	customerUsecase usecases.CustomerUsecase
}

func (h *CustomerHandler) CreateCustomer(c *gin.Context) interface{} {
	recipient := new(models.Recipient)
	err := c.BindJSON(recipient)
	if err != nil {
		return dtmcli.ErrFailure
	}
	transactionId := c.Query(gid)

	return h.customerUsecase.CreateCustomer(
		transactionId,
		recipient,
	)
}

func (h *CustomerHandler) CreateCustomerCompensate(c *gin.Context) interface{} {
	transactionId := c.Query(gid)

	return h.customerUsecase.CreateCustomerCompensate(
		transactionId,
	)
}
