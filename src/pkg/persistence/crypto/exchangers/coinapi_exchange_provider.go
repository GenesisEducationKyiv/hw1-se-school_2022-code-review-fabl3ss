package exchangers

import (
	"fmt"
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain/models"
	"genesis_test_case/src/pkg/domain/usecase"
	"genesis_test_case/src/pkg/utils"
	"net/http"
	"os"
)

type CoinApiProviderFactory struct{}

func (factory CoinApiProviderFactory) CreateExchangeProvider() usecase.ExchangeProvider {
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
}

func (c *coinApiExchangeProvider) GetCurrencyRate(pair *models.CurrencyPair) (*models.CurrencyRate, error) {
	resp, err := c.makeAPIRequest(pair)
	if err != nil {
		return nil, err
	}

	return resp.toDefaultRate(), nil
}

func (c *coinApiExchangeProvider) makeAPIRequest(pair *models.CurrencyPair) (*coinAPIExchangerResponse, error) {
	url := fmt.Sprintf(
		c.exchangeTemplateUrl,
		pair.GetBaseCurrency(),
		pair.GetQuoteCurrency(),
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

type coinAPIExchangerResponse struct {
	Time          string  `json:"time"`
	BaseCurrency  string  `json:"asset_id_base"`
	QuoteCurrency string  `json:"asset_id_quote"`
	Rate          float64 `json:"rate"`
}

func (c *coinAPIExchangerResponse) toDefaultRate() *models.CurrencyRate {
	return &models.CurrencyRate{
		Price: c.Rate,
		CurrencyPair: *models.NewCurrencyPair(
			c.BaseCurrency,
			c.QuoteCurrency,
		),
	}
}
