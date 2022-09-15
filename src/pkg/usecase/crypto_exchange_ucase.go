package usecase

import (
	"errors"
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/domain"
	myerr "genesis_test_case/src/pkg/types/errors"
)

type CryptoExchangerUsecase struct {
	pair           *domain.CurrencyPair
	cryptoProvider ExchangeProvider
	cache          CryptoCache
}

func NewCryptoExchangeUsecase(
	pair *domain.CurrencyPair,
	crypto ExchangeProvider,
	cache CryptoCache,
) http.CryptoExchangerUsecase {
	return &CryptoExchangerUsecase{
		pair:           pair,
		cryptoProvider: crypto,
		cache:          cache,
	}
}

func (c *CryptoExchangerUsecase) GetCurrentExchangePrice() (float64, error) {
	cacheRate, err := c.cache.GetCurrencyCache(config.EnvCryptoCacheKey)
	if err != nil {
		if errors.Is(err, myerr.ErrNoCache) {
			rate, err := c.cryptoProvider.GetCurrencyRate(c.pair)
			if err != nil {
				return -1, err
			}
			err = c.cache.SetCurrencyCache(config.EnvCryptoCacheKey, rate)
			if err != nil {
				return -1, err
			}
			return rate.Price, nil
		}

		return -1, err
	}

	return cacheRate.Price, nil
}
