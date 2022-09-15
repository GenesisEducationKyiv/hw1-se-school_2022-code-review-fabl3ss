package exchangers

import (
	"fmt"
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/usecase"
	"genesis_test_case/src/pkg/utils"
	"net/http"
	"os"
)

type coinAPIExchangerResponse struct {
	Time          string  `json:"time"`
	BaseCurrency  string  `json:"asset_id_base"`
	QuoteCurrency string  `json:"asset_id_quote"`
	Rate          float64 `json:"rate"`
}

func (c *coinAPIExchangerResponse) toDefaultRate() *domain.CurrencyRate {
	return &domain.CurrencyRate{
		Price: c.Rate,
		CurrencyPair: domain.CurrencyPair{
			BaseCurrency:  c.BaseCurrency,
			QuoteCurrency: c.QuoteCurrency,
		},
	}
}

type CoinApiProviderFactory struct{}

func (factory CoinApiProviderFactory) CreateExchangeProviderNode() usecase.ExchangeProviderNode {
	return &coinApiExchangeProvider{
		apiKey:              os.Getenv(config.EnvCoinAPIKey),
		apiKeyHeader:        "X-CoinAPI-Key",
		exchangeTemplateUrl: "https://rest.coinapi.io/v1/exchangerate/%v/%v",
	}
}

type coinApiExchangeProvider struct {
	apiKey              string
	apiKeyHeader        string
	exchangeTemplateUrl string
	next                usecase.ExchangeProviderNode
}

func (*coinApiExchangeProvider) GetWeekAverageChart(pair *domain.CurrencyPair) ([]float64, error) {
	panic("unimplemented")
}

func (c *coinApiExchangeProvider) SetNext(service usecase.ExchangeProviderNode) {
	c.next = service
}

func (c *coinApiExchangeProvider) GetCurrencyRate(pair *domain.CurrencyPair) (*domain.CurrencyRate, error) {
	rate, err := c.getExchangeRate(pair)
	if err != nil && c.next != nil {
		return c.next.GetCurrencyRate(pair)
	}

	return rate, err
}

func (c *coinApiExchangeProvider) getExchangeRate(pair *domain.CurrencyPair) (*domain.CurrencyRate, error) {
	resp, err := c.makeAPIRequest(pair)
	if err != nil {
		return nil, err
	}

	return resp.toDefaultRate(), nil
}

func (c *coinApiExchangeProvider) makeAPIRequest(pair *domain.CurrencyPair) (*coinAPIExchangerResponse, error) {
	url := fmt.Sprintf(
		c.exchangeTemplateUrl,
		pair.BaseCurrency,
		pair.QuoteCurrency,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(c.apiKeyHeader, c.apiKey)
	coinAPIRate := new(coinAPIExchangerResponse)
	err = utils.DoHttpAndParseBody(req, coinAPIRate)
	if err != nil {
		return nil, err
	}

	return coinAPIRate, nil
}
