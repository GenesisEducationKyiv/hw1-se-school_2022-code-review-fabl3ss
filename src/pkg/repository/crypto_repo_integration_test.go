package repository

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/stretchr/testify/require"
)

func TestGetCurrencyRate(t *testing.T) {
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	cryptoRepo := NewCryptoRepository()
	baseCurrency := "BTC"
	quoteCurrency := "UAH"

	rate, err := cryptoRepo.GetCurrencyRate(baseCurrency, quoteCurrency)
	require.NoError(t, err)
	require.Equal(t, rate.BaseCurrency, baseCurrency)
	require.Equal(t, rate.QuoteCurrency, quoteCurrency)
	require.NotEmpty(t, rate.Price)
}

func TestGetWeekChart(t *testing.T) {
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	cryptoRepo := NewCryptoRepository()

	candles, err := cryptoRepo.GetWeekChart()
	require.NoError(t, err)
	require.NotEmpty(t, candles)
}
