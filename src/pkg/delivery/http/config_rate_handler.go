package http

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain/models"
	"github.com/gofiber/fiber/v2"
	"os"
)

type ConfigRateHandler struct {
	exchangeUsecase CryptoExchangerUsecase
	presenter       ResponsePresenter
}

func NewConfigRateHandler(exchanger CryptoExchangerUsecase, presenter ResponsePresenter) *ConfigRateHandler {
	return &ConfigRateHandler{
		exchangeUsecase: exchanger,
		presenter:       presenter,
	}
}

func (r *ConfigRateHandler) GetCurrencyRate(c *fiber.Ctx) error {
	defaultRate := models.NewCurrencyPair(
		os.Getenv(config.EnvBaseCurrency),
		os.Getenv(config.EnvQuoteCurrency),
	)

	rate, err := r.exchangeUsecase.GetCurrentExchangePrice(defaultRate)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return r.presenter.PresentExchangeRate(c,
		&RateResponse{
			Rate: rate,
		},
	)
}
