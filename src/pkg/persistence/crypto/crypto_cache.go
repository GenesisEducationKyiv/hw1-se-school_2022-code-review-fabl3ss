package crypto

import (
	"encoding/json"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/usecase"
)

type cryptoCache struct {
	cacheProvider usecase.Cache
}

func NewCryptoCache(cache usecase.Cache) usecase.CryptoCache {
	return &cryptoCache{
		cacheProvider: cache,
	}
}

func (c *cryptoCache) GetCurrencyCache(key string) (*domain.CurrencyRate, error) {
	rateByte, err := c.cacheProvider.GetCache(key)
	if err != nil {
		return nil, err
	}

	rate := new(domain.CurrencyRate)
	err = json.Unmarshal(rateByte, rate)
	if err != nil {
		return nil, err
	}

	return rate, nil
}

func (c *cryptoCache) SetCurrencyCache(key string, rate *domain.CurrencyRate) error {
	return c.cacheProvider.SetCache(key, rate)
}
