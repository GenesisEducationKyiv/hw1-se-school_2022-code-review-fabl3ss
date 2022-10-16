package usecases

import "customers/src/pkg/domain/models"

type CustomerUsecase interface {
	CreateCustomer(transactionId string, recipient *models.Recipient) error
	CreateCustomerCompensate(transactionId string) error
}
