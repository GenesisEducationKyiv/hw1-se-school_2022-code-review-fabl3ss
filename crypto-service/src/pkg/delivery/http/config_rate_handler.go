package http

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/delivery/http/responses"
	"genesis_test_case/src/pkg/domain/models"
	"genesis_test_case/src/pkg/domain/usecases"
	"os"

	"github.com/gofiber/fiber/v2"
)

type ConfigRateHandler struct {
	exchangeUsecase usecases.CryptoExchangerUsecase
	presenter       ResponsePresenter
}

func NewConfigRateHandler(exchanger usecases.CryptoExchangerUsecase, presenter ResponsePresenter) *ConfigRateHandler {
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
		&responses.RateResponse{
			Rate: rate,
		},
	)
}
