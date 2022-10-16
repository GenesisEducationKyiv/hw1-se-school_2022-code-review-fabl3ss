package usecase

import (
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/models"
	"genesis_test_case/src/pkg/domain/usecases"
)

type subscriptionCustomerUsecase struct {
	subscriptionUsecase usecases.SubscriptionUsecase
	customerProvider    application.CustomerProvider
}

func NewSubscriptionCustomerUsecase(
	subscription usecases.SubscriptionUsecase,
	customer application.CustomerProvider,
) usecases.SubscriptionUsecase {
	return &subscriptionCustomerUsecase{
		subscriptionUsecase: subscription,
		customerProvider:    customer,
	}
}

func (s *subscriptionCustomerUsecase) Subscribe(recipient *models.Recipient) error {
	err := s.subscriptionUsecase.Subscribe(recipient)
	if err != nil {
		return err
	}

	return s.customerProvider.CreateCustomer(recipient)
}
