package domain

import "genesis_test_case/src/config"

type CurrencyRate struct {
	Price         string `json:"amount"`
	BaseCurrency  string `json:"base"`
	QuoteCurrency string `json:"currency"`
}

type CandleProps struct {
	Base        string
	Granularity string
	Start       string
	End         string
}

type CryptoUsecase interface {
	GetConfigCurrencyRate(cfg *config.Config) (int, error)
}

type CryptoRepository interface {
	GetWeekChart() ([]float64, error)
	GetCandles(candleProps *CandleProps) ([][]float64, error)
	GetCurrencyRate(base string, quoted string) (*CurrencyRate, error)
}
