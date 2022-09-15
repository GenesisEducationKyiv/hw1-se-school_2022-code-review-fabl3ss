package exchangers

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/require"
)

func TestGetCurrencyRate(t *testing.T) {
	if err := godotenv.Load("../../../../../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	coinbaseRepo := NewcoinbaseExchangeProvidersitory(
		os.Getenv(config.EnvCoinbaseApiExchangeUrl),
		os.Getenv(config.EnvCoinbaseApiCandlesUrl),
	)
	pair := &domain.CurrencyPair{
		BaseCurrency:  "BTC",
		QuoteCurrency: "UAH",
	}

	rate, err := coinbaseRepo.GetCurrencyRate(pair)
	require.NoError(t, err)
	require.Equal(t, pair.BaseCurrency, rate.BaseCurrency)
	require.Equal(t, pair.QuoteCurrency, rate.QuoteCurrency)
	require.NotEmpty(t, rate.Price)
}

func TestGetWeekChart(t *testing.T) {
	if err := godotenv.Load("../../../../../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	coinbaseRepo := NewcoinbaseExchangeProvidersitory(
		os.Getenv(config.EnvCoinbaseApiExchangeUrl),
		os.Getenv(config.EnvCoinbaseApiCandlesUrl),
	)
	pair := &domain.CurrencyPair{
		BaseCurrency:  "BTC",
		QuoteCurrency: "UAH",
	}

	candles, err := coinbaseRepo.GetWeekAverageChart(pair)
	require.NoError(t, err)
	require.NotEmpty(t, candles)
}
