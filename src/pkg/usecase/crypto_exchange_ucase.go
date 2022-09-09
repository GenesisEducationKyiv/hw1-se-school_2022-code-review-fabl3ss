package usecase

import (
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/domain"
)

type CryptoExchangerUsecase struct {
	pair           *domain.CurrencyPair
	cryptoProvider CryptoRepository
}

func NewCryptoExchangeUsecase(
	pair *domain.CurrencyPair,
	crypto CryptoRepository,
) http.CryptoExchangerUsecase {
	return &CryptoExchangerUsecase{
		pair:           pair,
		cryptoProvider: crypto,
	}
}

func (c *CryptoExchangerUsecase) GetCurrentExchangePrice() (float64, error) {
	rate, err := c.cryptoProvider.GetCurrencyRate(c.pair)
	if err != nil {
		return 0, err
	}

	return rate.Price, nil
}
