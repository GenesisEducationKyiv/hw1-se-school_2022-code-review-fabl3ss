package application

import "customers/src/pkg/domain/models"

type CustomerRepository interface {
	CreateCustomer(order *models.Customer) error
	UpdateCustomerStatus(transactionId, status string) error
}
