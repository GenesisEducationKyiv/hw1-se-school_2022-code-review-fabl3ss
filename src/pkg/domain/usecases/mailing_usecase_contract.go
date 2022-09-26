package usecases

import "genesis_test_case/src/pkg/domain/models"

type CryptoMailingUsecase interface {
	SendCurrencyRate() ([]models.EmailAddress, error)
}
