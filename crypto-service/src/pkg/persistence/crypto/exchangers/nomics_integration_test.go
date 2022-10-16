package exchangers

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestNomicsGetCurrencyRate(t *testing.T) {
	if err := godotenv.Load("../../../../../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	nomicsProvider := NomicsProviderFactory{}.CreateExchangeProvider()

	GetCurrencyRateTest(nomicsProvider, t)
}
