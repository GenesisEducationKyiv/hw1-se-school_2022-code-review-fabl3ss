package charts

import (
	"genesis_test_case/src/pkg/domain/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetWeekChart(t *testing.T) {
	coinbaseRepo := CoinbaseProviderFactory{}.CreateChartProvider()
	pair := models.NewCurrencyPair(
		"BTC",
		"UAH",
	)

	candles, err := coinbaseRepo.GetWeekAverageChart(pair)
	require.NoError(t, err)
	require.NotEmpty(t, candles)
}
