package application

import (
	"orders/src/pkg/domain/models"
	"orders/src/pkg/domain/usecases"

	"github.com/dtm-labs/client/dtmcli"
)

func NewOrdersUsecase(coordinatorURL, customerURL string) usecases.OrdersUsecase {
	return &ordersUsecase{
		coordinatorURL:     coordinatorURL,
		customerServiceURL: customerURL,
	}
}

type ordersUsecase struct {
	coordinatorURL     string
	customerServiceURL string
}

func (o *ordersUsecase) RegisterCustomer(recipient *models.Recipient) (string, error) {
	gid := dtmcli.MustGenGid(o.coordinatorURL)
	saga := dtmcli.NewSaga(o.coordinatorURL, gid).
		Add(
			o.customerServiceURL+"/create-customer",
			o.customerServiceURL+"/create-customer-compensate", *recipient,
		)

	return gid, saga.Submit()
}
