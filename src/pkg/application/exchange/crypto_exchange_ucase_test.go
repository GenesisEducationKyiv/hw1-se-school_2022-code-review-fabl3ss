package usecase

import (
	"genesis_test_case/src/config"
	mocks "genesis_test_case/src/pkg/application/mocks"
	"genesis_test_case/src/pkg/domain/models"
	myerr "genesis_test_case/src/pkg/types/errors"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"
)

func TestGetConfigCurrencyRate(t *testing.T) {
	if err := godotenv.Load("../../../../.env"); err != nil {
		t.Error("Error loading .env file")
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mockExchanger := mocks.NewMockExchangeProvider(ctl)
	mockCryptoCache := mocks.NewMockCryptoCache(ctl)
	BTCUAHPair := models.NewCurrencyPair(
		os.Getenv(config.EnvBaseCurrency),
		os.Getenv(config.EnvQuoteCurrency),
	)

	cryptoExchangeUsecase := NewCryptoExchangeUsecase(
		mockExchanger,
		mockCryptoCache,
	)

	mockResp := &models.CurrencyRate{
		CurrencyPair: *BTCUAHPair,
		Price:        123.123,
	}
	mockCryptoCache.EXPECT().GetCurrencyCache(config.CryptoCacheKey).Return(nil, myerr.ErrNoCache)
	mockExchanger.EXPECT().GetCurrencyRate(BTCUAHPair).Return(mockResp, nil)
	mockCryptoCache.EXPECT().SetCurrencyCache(config.CryptoCacheKey, mockResp).Return(nil)
	_, err := cryptoExchangeUsecase.GetCurrentExchangePrice(BTCUAHPair)

	require.NoError(t, err)
}
