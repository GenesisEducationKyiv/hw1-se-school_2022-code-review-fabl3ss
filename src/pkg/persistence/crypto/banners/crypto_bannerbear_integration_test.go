package banners

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestGetCryptoBannerUrl(t *testing.T) {
	if err := godotenv.Load("../../../../../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	bannerRepo := NewCryptoBannerBearRepository(
		os.Getenv(config.EnvBannerApiToken),
		os.Getenv(config.EnvBannerApiUrl),
		os.Getenv(config.EnvCryptoBannerTemplate),
	)
	chart := []float64{
		0.1,
		0.2,
		0.3,
		0.4,
	}
	rate := &domain.CurrencyRate{
		CurrencyPair: domain.CurrencyPair{
			BaseCurrency:  "BTC",
			QuoteCurrency: "UAH",
		},
		Price: 123.123,
	}

	jpgUrl, err := bannerRepo.GetCryptoBannerUrl(chart, rate)
	require.NoError(t, err)
	require.NotEmpty(t, jpgUrl)
}
