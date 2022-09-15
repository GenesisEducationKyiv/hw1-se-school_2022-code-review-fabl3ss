package charts

import (
	"fmt"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/usecase"
	"genesis_test_case/src/pkg/utils"
	"strconv"
	"time"
)

type CoinbaseProviderFactory struct{}

func (factory CoinbaseProviderFactory) CreateChartProvider() usecase.ChartProvider {
	return &coinbaseChartProvider{
		chartTemplateUrl: "https://rest.coinapi.io/v1/exchangerate/%v/%v",
	}
}

type coinbaseChartProvider struct {
	chartTemplateUrl string
}

type chartProps struct {
	Base        string
	Quote       string
	Granularity string
	Start       string
	End         string
}

func (c *coinbaseChartProvider) GetWeekAverageChart(pair *domain.CurrencyPair) ([]float64, error) {
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

func (c *coinbaseChartProvider) getWeekCandles(pair *domain.CurrencyPair) ([][]float64, error) {
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

func (c *coinbaseChartProvider) getChart(candleProps *chartProps) ([][]float64, error) {
	var candles [][]float64
	url := fmt.Sprintf(
		c.chartTemplateUrl,
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
