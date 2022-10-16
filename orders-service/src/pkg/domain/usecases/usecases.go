package usecases

import "orders/src/pkg/domain/models"

type OrdersUsecase interface {
	RegisterCustomer(recipient *models.Recipient) (string, error)
}
