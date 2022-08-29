package repository

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/utils"
)

type cryptoRepository struct{}

func NewCryptoRepository() domain.CryptoRepository {
	return &cryptoRepository{}
}

func (c *cryptoRepository) GetCandles(candleProps *domain.CandleProps) ([][]float64, error) {
	cfg := config.Get()
	var candles [][]float64
	url := fmt.Sprintf(
		cfg.CryptoApiCandlesUrl,
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

func (c *cryptoRepository) GetCurrencyRate(base string, quoted string) (*domain.CurrencyRate, error) {
	cfg := config.Get()
	url := fmt.Sprintf(
		cfg.CryptoApiFormatUrl,
		base,
		quoted,
	)
	rate := struct {
		Rate domain.CurrencyRate `json:"data"`
	}{}

	err := utils.GetAndParseBody(url, &rate)
	if err != nil {
		return nil, err
	}
	rate.Rate.Price = strings.Split(rate.Rate.Price, ".")[0]

	return &rate.Rate, nil
}

func (c *cryptoRepository) GetWeekChart() ([]float64, error) {
	var averageCandles []float64
	weekCandles, err := c.getWeekCandles()
	if err != nil {
		return nil, err
	}

	for i := len(weekCandles) - 1; i >= 0; i-- {
		// [i][3] -> opening price (first trade) in the bucket interval
		averageCandles = append(averageCandles, weekCandles[i][3])
	}

	return averageCandles, nil
}

func (c *cryptoRepository) getWeekCandles() ([][]float64, error) {
	cfg := config.Get()
	nowUtc := time.Now().UTC()
	weekCandlesProps := &domain.CandleProps{
		Base:        cfg.BaseCurrency,
		Granularity: strconv.Itoa(int(time.Hour.Seconds())),
		Start:       nowUtc.AddDate(0, 0, -7).Format(time.RFC3339),
		End:         nowUtc.Format(time.RFC3339),
	}

	candles, err := c.GetCandles(weekCandlesProps)
	if err != nil {
		return nil, err
	}

	return candles, nil
}
