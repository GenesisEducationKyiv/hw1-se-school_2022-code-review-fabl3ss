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

type nomicsExchangerResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (c *nomicsExchangerResponse) toDefaultRate(quote string) (*domain.CurrencyRate, error) {
	floatPrice, err := utils.StringToFloat64(c.Price)
	if err != nil {
		return nil, err
	}
	return &domain.CurrencyRate{
		Price: floatPrice,
		CurrencyPair: domain.CurrencyPair{
			BaseCurrency:  c.Symbol,
			QuoteCurrency: quote,
		},
	}, nil
}

type NomicsProviderFactory struct{}

func (factory NomicsProviderFactory) CreateRateService() usecase.ExchangeProviderNode {
	return &nomicsExchangeProvider{
		exchangeTemplateUrl: "https://api.nomics.com/v1/currencies/ticker?key=%v&ids=%v&interval=1d&convert=%v",
		apiKey:              os.Getenv(config.EnvNomicsApiKey),
	}
}

type nomicsExchangeProvider struct {
	exchangeTemplateUrl string
	apiKey              string
	next                usecase.ExchangeProviderNode
}

func (*nomicsExchangeProvider) GetWeekAverageChart(pair *domain.CurrencyPair) ([]float64, error) {
	panic("unimplemented")
}

func (n *nomicsExchangeProvider) SetNext(service usecase.ExchangeProviderNode) {
	n.next = service
}

func (n *nomicsExchangeProvider) GetCurrencyRate(pair *domain.CurrencyPair) (*domain.CurrencyRate, error) {
	rate, err := n.getExchangeRate(pair)
	if err != nil && n.next != nil {
		return (n.next).GetCurrencyRate(pair)
	}

	return rate, err
}

func (n *nomicsExchangeProvider) getExchangeRate(pair *domain.CurrencyPair) (*domain.CurrencyRate, error) {
	resp, err := n.makeAPIRequest(pair)
	if err != nil {
		return nil, err
	}

	return resp.toDefaultRate(pair.QuoteCurrency)
}

func (n *nomicsExchangeProvider) makeAPIRequest(pair *domain.CurrencyPair) (*nomicsExchangerResponse, error) {
	url := fmt.Sprintf(
		n.exchangeTemplateUrl,
		n.exchangeTemplateUrl,
		pair.BaseCurrency,
		pair.QuoteCurrency,
	)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var nomicsRate []nomicsExchangerResponse
	err = utils.DoHttpAndParseBody(req, &nomicsRate)
	if err != nil {
		return nil, err
	}
	return &nomicsRate[0], nil
}
