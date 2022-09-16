package usecase

import (
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/domain"
	myerr "genesis_test_case/src/pkg/types/errors"
	"genesis_test_case/src/pkg/usecase"
)

type subscriptionUsecase struct {
	storage usecase.EmailStorage
}

func NewSubscriptionUsecase(storage usecase.EmailStorage) http.SubscriptionUsecase {
	return &subscriptionUsecase{
		storage: storage,
	}
}

func (s *subscriptionUsecase) Subscribe(recipient *domain.Recipient) error {
	if recipient == nil {
		return myerr.ErrNoDataProvided
	}
	err := s.storage.AddEmail(recipient.Email)
	if err != nil {
		return err
	}

	return nil
}
