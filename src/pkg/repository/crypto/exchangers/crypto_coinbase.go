package exchangers

import (
	"fmt"
	"strconv"
	"time"

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
	floatPrice, err := strconv.ParseFloat(c.Amount, 64)
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

type chartProps struct {
	Base        string
	Quote       string
	Granularity string
	Start       string
	End         string
}

type cryptoCoinbaseRepo struct {
	ExchangeEndpoint string
	ChartEndpoint    string
}

func NewCryptoCoinbaseRepository(exchangeEndpoint string, chartEndpoint string) usecase.CryptoRepository {
	return &cryptoCoinbaseRepo{
		ExchangeEndpoint: exchangeEndpoint,
		ChartEndpoint:    chartEndpoint,
	}
}

func (c *cryptoCoinbaseRepo) GetCurrencyRate(pair *domain.CurrencyPair) (*domain.CurrencyRate, error) {
	url := fmt.Sprintf(
		c.ExchangeEndpoint,
		pair.BaseCurrency,
		pair.QuoteCurrency,
	)
	rate := new(coinbaseExchangerResponse)

	err := utils.GetAndParseBody(url, rate)
	if err != nil {
		return nil, err
	}

	return rate.toDefaultRate()
}

func (c *cryptoCoinbaseRepo) GetWeekAverageChart(pair *domain.CurrencyPair) ([]float64, error) {
	var averageCandles []float64
	weekCandles, err := c.getWeekCandles(pair)
	if err != nil {
		return nil, err
	}

	for i := len(weekCandles) - 1; i >= 0; i-- {
		// [i][3] -> opening price (first trade) in the bucket interval
		averageCandles = append(averageCandles, weekCandles[i][3])
	}

	return averageCandles, nil
}

func (c *cryptoCoinbaseRepo) getWeekCandles(pair *domain.CurrencyPair) ([][]float64, error) {
	nowUtc := time.Now().UTC()
	weekCandlesProps := &chartProps{
		Base:        pair.BaseCurrency,
		Quote:       pair.QuoteCurrency,
		Granularity: strconv.Itoa(int(time.Hour.Seconds())),
		Start:       nowUtc.AddDate(0, 0, -7).Format(time.RFC3339),
		End:         nowUtc.Format(time.RFC3339),
	}

	candles, err := c.getChart(weekCandlesProps)
	if err != nil {
		return nil, err
	}

	return candles, nil
}

func (c *cryptoCoinbaseRepo) getChart(candleProps *chartProps) ([][]float64, error) {
	var candles [][]float64
	url := fmt.Sprintf(
		c.ChartEndpoint,
		candleProps.Base,
		candleProps.Granularity,
		candleProps.Start,
		candleProps.End,
	)

	err := utils.GetAndParseBody(url, &candles)
	if err != nil {
		return nil, err
	}

	return candles, nil
}
