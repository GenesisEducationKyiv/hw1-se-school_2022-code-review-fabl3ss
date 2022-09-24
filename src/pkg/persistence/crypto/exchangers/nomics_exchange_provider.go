package exchangers

import (
	"fmt"
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/models"
	"genesis_test_case/src/pkg/utils"
	"net/http"
	"os"
)

type NomicsProviderFactory struct{}

func (factory NomicsProviderFactory) CreateExchangeProvider() application.ExchangeProvider {
	return &nomicsExchangeProvider{
		exchangeTemplateUrl: "https://api.nomics.com/v1/currencies/ticker?key=%v&ids=%v&interval=1d&convert=%v",
		apiKey:              os.Getenv(config.EnvNomicsApiKey),
	}
}

type nomicsExchangeProvider struct {
	exchangeTemplateUrl string
	apiKey              string
}

func (n *nomicsExchangeProvider) GetCurrencyRate(pair *models.CurrencyPair) (*models.CurrencyRate, error) {
	resp, err := n.makeAPIRequest(pair)
	if err != nil {
		return nil, err
	}

	return resp.toDefaultRate(pair.GetQuoteCurrency())
}

func (n *nomicsExchangeProvider) makeAPIRequest(pair *models.CurrencyPair) (*nomicsExchangerResponse, error) {
	url := fmt.Sprintf(
		n.exchangeTemplateUrl,
		n.apiKey,
		pair.GetBaseCurrency(),
		pair.GetQuoteCurrency(),
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

type nomicsExchangerResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (c *nomicsExchangerResponse) toDefaultRate(quote string) (*models.CurrencyRate, error) {
	floatPrice, err := utils.StringToFloat64(c.Price)
	if err != nil {
		return nil, err
	}
	return &models.CurrencyRate{
		Price: floatPrice,
		CurrencyPair: *models.NewCurrencyPair(
			c.Symbol,
			quote,
		),
	}, nil
}
