package exchangers

import (
	"testing"
)

func TestCoinbaseGetCurrencyRate(t *testing.T) {
	coinbaseProvider := CoinbaseProviderFactory{}.CreateExchangeProvider()
	GetCurrencyRateTest(coinbaseProvider, t)
}
