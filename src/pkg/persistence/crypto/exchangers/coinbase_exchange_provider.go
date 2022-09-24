package exchangers

import (
	"fmt"
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/models"
	"genesis_test_case/src/pkg/utils"
)

type CoinbaseProviderFactory struct{}

func (factory CoinbaseProviderFactory) CreateExchangeProvider() application.ExchangeProvider {
	return &coinbaseExchangeProvider{
		exchangeTemplateUrl: "https://api.coinbase.com/v2/prices/%s-%s/spot",
	}
}

type coinbaseExchangeProvider struct {
	exchangeTemplateUrl string
}

func (c *coinbaseExchangeProvider) GetCurrencyRate(pair *models.CurrencyPair) (*models.CurrencyRate, error) {
	resp, err := c.makeAPIRequest(pair)
	if err != nil {
		return nil, err
	}

	return resp.toDefaultRate()
}

func (c *coinbaseExchangeProvider) makeAPIRequest(pair *models.CurrencyPair) (*coinbaseExchangerResponse, error) {
	url := fmt.Sprintf(
		c.exchangeTemplateUrl,
		pair.GetBaseCurrency(),
		pair.GetQuoteCurrency(),
	)
	rate := new(coinbaseExchangerResponse)

	err := utils.GetAndParseBody(url, rate)
	if err != nil {
		return nil, err
	}

	return rate, nil
}

type coinbaseCurrencyRate struct {
	Amount        string `json:"amount"`
	BaseCurrency  string `json:"base"`
	QuoteCurrency string `json:"currency"`
}

type coinbaseExchangerResponse struct {
	coinbaseCurrencyRate `json:"data"`
}

func (c *coinbaseCurrencyRate) toDefaultRate() (*models.CurrencyRate, error) {
	floatPrice, err := utils.StringToFloat64(c.Amount)
	if err != nil {
		return nil, err
	}

	return &models.CurrencyRate{
		Price: floatPrice,
		CurrencyPair: *models.NewCurrencyPair(
			c.BaseCurrency,
			c.QuoteCurrency,
		),
	}, nil
}
