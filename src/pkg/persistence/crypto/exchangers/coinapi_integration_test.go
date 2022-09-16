package exchangers

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestCoinApiGetCurrencyRate(t *testing.T) {
	if err := godotenv.Load("../../../../../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	coinApiProvider := CoinApiProviderFactory{}.CreateExchangeProvider()

	GetCurrencyRateTest(coinApiProvider, t)
}
