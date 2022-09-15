package exchangers

import (
	"fmt"

	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/usecase"
	"genesis_test_case/src/pkg/utils"
)

type coinbaseCurrencyRate struct {
	Amount        string `json:"amount"`
	BaseCurrency  string `json:"base"`
	QuoteCurrency string `json:"currency"`
}

type coinbaseExchangerResponse struct {
	coinbaseCurrencyRate `json:"data"`
}

func (c *coinbaseCurrencyRate) toDefaultRate() (*domain.CurrencyRate, error) {
	floatPrice, err := utils.StringToFloat64(c.Amount)
	if err != nil {
		return nil, err
	}

	return &domain.CurrencyRate{
		Price: floatPrice,
		CurrencyPair: domain.CurrencyPair{
			BaseCurrency:  c.BaseCurrency,
			QuoteCurrency: c.QuoteCurrency,
		},
	}, nil
}

type CoinbaseProviderFactory struct{}

func (factory CoinbaseProviderFactory) CreateExchangeProvider() usecase.ExchangeProvider {
	return &coinbaseExchangeProvider{
		exchangeTemplateUrl: "https://rest.coinapi.io/v1/exchangerate/%v/%v",
	}
}

type coinbaseExchangeProvider struct {
	exchangeTemplateUrl string
}

func (c *coinbaseExchangeProvider) GetCurrencyRate(pair *domain.CurrencyPair) (*domain.CurrencyRate, error) {
	resp, err := c.makeAPIRequest(pair)
	if err != nil {
		return nil, err
	}

	return resp.toDefaultRate()
}

func (c *coinbaseExchangeProvider) makeAPIRequest(pair *domain.CurrencyPair) (*coinbaseExchangerResponse, error) {
	url := fmt.Sprintf(
		c.exchangeTemplateUrl,
		pair.BaseCurrency,
		pair.QuoteCurrency,
	)
	rate := new(coinbaseExchangerResponse)

	err := utils.GetAndParseBody(url, rate)
	if err != nil {
		return nil, err
	}

	return rate, nil
}
