package repository

import (
	"genesis_test_case/src/pkg/domain"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestGetCryptoBannerUrl(t *testing.T) {
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	imageRepo := NewImageRepository()
	chart := []float64{
		0.1,
		0.2,
		0.3,
		0.4,
	}
	rate := &domain.CurrencyRate{
		BaseCurrency:  "BTC",
		QuoteCurrency: "UAH",
		Price:         "12345",
	}

	jpgUrl, err := imageRepo.GetCryptoBannerUrl(chart, rate)
	require.NoError(t, err)
	require.NotEmpty(t, jpgUrl)
}
