package usecase

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain"
	mocks "genesis_test_case/src/pkg/domain/mocks"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestGetConfigCurrencyRate(t *testing.T) {
	if err := godotenv.Load("../../../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockCryptoRepo := mocks.NewMockCryptoRepository(ctl)
	BTCUAHPair := &domain.CurrencyPair{
		BaseCurrency:  os.Getenv(config.EnvBaseCurrency),
		QuoteCurrency: os.Getenv(config.EnvQuoteCurrency),
	}

	cryptoExchangeUsecase := NewCryptoExchangeUsecase(
		BTCUAHPair,
		mockCryptoRepo,
	)

	mockResp := &domain.CurrencyRate{
		CurrencyPair: *BTCUAHPair,
		Price:        123.123,
	}

	mockCryptoRepo.EXPECT().GetCurrencyRate(BTCUAHPair).Return(mockResp, nil)
	_, err := cryptoExchangeUsecase.GetCurrentExchangePrice()
	require.NoError(t, err)
}
