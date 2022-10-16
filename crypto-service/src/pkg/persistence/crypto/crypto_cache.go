package crypto

import (
	"encoding/json"
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/models"
)

type cryptoCache struct {
	cacheProvider application.Cache
}

func NewCryptoCache(cache application.Cache) application.CryptoCache {
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
