package http

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain"
	"github.com/gofiber/fiber/v2"
	"os"
)

type ConfigRateHandler struct {
	exchangeUsecase CryptoExchangerUsecase
}

func NewConfigRateHandler(exchanger CryptoExchangerUsecase) *ConfigRateHandler {
	return &ConfigRateHandler{
		exchangeUsecase: exchanger,
	}
}

func (r *ConfigRateHandler) GetCurrencyRate(c *fiber.Ctx) error {
	defaultRate := domain.NewCurrencyPair(
		os.Getenv(config.EnvBaseCurrency),
		os.Getenv(config.EnvQuoteCurrency),
	)

	rate, err := r.exchangeUsecase.GetCurrentExchangePrice(defaultRate)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(rate)
}
