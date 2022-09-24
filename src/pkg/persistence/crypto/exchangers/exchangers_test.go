package exchangers

import (
	"genesis_test_case/src/pkg/domain/models"
	"genesis_test_case/src/pkg/domain/usecase"
	"testing"

	"github.com/stretchr/testify/require"
)

func GetCurrencyRateTest(exchanger usecase.ExchangeProvider, t *testing.T) {
	pair := models.NewCurrencyPair(
		"BTC",
		"UAH",
	)

	rate, err := exchanger.GetCurrencyRate(pair)

	require.NoError(t, err)
	require.Equal(t, pair.GetBaseCurrency(), rate.GetBaseCurrency())
	require.Equal(t, pair.GetQuoteCurrency(), rate.GetQuoteCurrency())
	require.NotEmpty(t, rate.Price)
}
