package usecase

import (
	"errors"
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/domain"
	myerr "genesis_test_case/src/pkg/types/errors"
	"genesis_test_case/src/pkg/usecase"
)

type CryptoExchangerUsecase struct {
	pair           *domain.CurrencyPair
	cryptoProvider usecase.ExchangeProvider
	cache          usecase.CryptoCache
}

func NewCryptoExchangeUsecase(
	crypto usecase.ExchangeProvider,
	cache usecase.CryptoCache,
) http.CryptoExchangerUsecase {
	return &CryptoExchangerUsecase{
		cryptoProvider: crypto,
		cache:          cache,
	}
}

func (c *CryptoExchangerUsecase) GetCurrentExchangePrice(pair *domain.CurrencyPair) (float64, error) {
	cacheRate, err := c.cache.GetCurrencyCache(config.CryptoCacheKey)
	if err != nil {
		if errors.Is(err, myerr.ErrNoCache) {
			rate, err := c.cryptoProvider.GetCurrencyRate(pair)
			if err != nil {
				return -1, err
			}
			err = c.cache.SetCurrencyCache(config.CryptoCacheKey, rate)
			if err != nil {
				return -1, err
			}
			return rate.Price, nil
		}

		return -1, err
	}

	return cacheRate.Price, nil
}
