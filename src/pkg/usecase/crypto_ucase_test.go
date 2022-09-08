package usecase

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain"
	mocks "genesis_test_case/src/pkg/domain/mocks"
	"genesis_test_case/src/pkg/repository"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetConfigCurrencyRate(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockCryptoRepo := mocks.NewMockCryptoRepository(ctl)
	repoCollection := &repository.Repositories{
		Crypto: mockCryptoRepo,
	}
	cryptoUsecase := NewCryptoUsecase(repoCollection)
	cfg := &config.Config{
		BaseCurrency:  "BTC",
		QuoteCurrency: "UAH",
	}
	mockResp := &domain.CurrencyRate{
		Price:         "12345",
		BaseCurrency:  cfg.BaseCurrency,
		QuoteCurrency: cfg.QuoteCurrency,
	}

	mockCryptoRepo.EXPECT().GetCurrencyRate(cfg.BaseCurrency, cfg.QuoteCurrency).Return(mockResp, nil)
	rate, err := cryptoUsecase.GetConfigCurrencyRate(cfg)
	require.NoError(t, err)

	rateInt, err := strconv.Atoi(mockResp.Price)
	require.NoError(t, err)
	require.Equal(t, rateInt, rate)
}
