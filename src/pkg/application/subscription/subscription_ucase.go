package usecase

import (
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/models"
	"genesis_test_case/src/pkg/domain/usecases"
	myerr "genesis_test_case/src/pkg/types/errors"
)

type subscriptionUsecase struct {
	storage application.EmailStorage
}

func NewSubscriptionUsecase(storage application.EmailStorage) usecases.SubscriptionUsecase {
	return &subscriptionUsecase{
		storage: storage,
	}
}

func (s *subscriptionUsecase) Subscribe(recipient *models.Recipient) error {
	if recipient == nil {
		return myerr.ErrNoDataProvided
	}
	err := s.storage.AddEmail(recipient.Email)
	if err != nil {
		return err
	}

	return nil
}
