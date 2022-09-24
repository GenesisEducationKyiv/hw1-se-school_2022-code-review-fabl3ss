package usecases

import "genesis_test_case/src/pkg/domain/models"

type CryptoExchangerUsecase interface {
	GetCurrentExchangePrice(pair *models.CurrencyPair) (float64, error)
}
