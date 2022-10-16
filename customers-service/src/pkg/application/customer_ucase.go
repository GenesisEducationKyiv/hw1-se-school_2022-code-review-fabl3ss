package application

import (
	"customers/src/pkg/domain/models"
	"customers/src/pkg/domain/usecases"
)

const (
	failedStatus  string = "Failed"
	createdStatus string = "Created"
)

func NewCustomerUsecase(customerRepo CustomerRepository) usecases.CustomerUsecase {
	return &customerUsecase{
		customerRepo: customerRepo,
	}
}

type customerUsecase struct {
	customerRepo CustomerRepository
}

func (c *customerUsecase) CreateCustomer(transactionId string, recipient *models.Recipient) error {
	order := &models.Customer{
		TransactionId: transactionId,
		Recipient:     *recipient,
		Status:        createdStatus,
	}

	err := c.customerRepo.CreateCustomer(order)
	if err != nil {
		return err
	}

	return err
}

func (c *customerUsecase) CreateCustomerCompensate(transactionId string) error {
	err := c.customerRepo.UpdateCustomerStatus(
		transactionId,
		failedStatus,
	)
	if err != nil {
		return err
	}
	return nil
}
