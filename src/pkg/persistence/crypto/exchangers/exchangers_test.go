package exchangers

import (
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/usecase"
	"testing"

	"github.com/stretchr/testify/require"
)

func GetCurrencyRateTest(exchanger usecase.ExchangeProvider, t *testing.T) {
	pair := &domain.CurrencyPair{
		BaseCurrency:  "BTC",
		QuoteCurrency: "UAH",
	}

	rate, err := exchanger.GetCurrencyRate(pair)

	require.NoError(t, err)
	require.Equal(t, pair.BaseCurrency, rate.BaseCurrency)
	require.Equal(t, pair.QuoteCurrency, rate.QuoteCurrency)
	require.NotEmpty(t, rate.Price)
}
