package crypto

import (
	"encoding/json"
	"genesis_test_case/src/pkg/domain/models"
	"genesis_test_case/src/pkg/domain/usecase"
)

type cryptoCache struct {
	cacheProvider usecase.Cache
}

func NewCryptoCache(cache usecase.Cache) usecase.CryptoCache {
	return &cryptoCache{
		cacheProvider: cache,
	}
}

func (c *cryptoCache) GetCurrencyCache(key string) (*models.CurrencyRate, error) {
	rateByte, err := c.cacheProvider.GetCache(key)
	if err != nil {
		return nil, err
	}

	rate := new(models.CurrencyRate)
	err = json.Unmarshal(rateByte, rate)
	if err != nil {
		return nil, err
	}

	return rate, nil
}

func (c *cryptoCache) SetCurrencyCache(key string, rate *models.CurrencyRate) error {
	return c.cacheProvider.SetCache(key, rate)
}
