package http

import (
	"genesis_test_case/src/pkg/domain/models"
	"github.com/gofiber/fiber/v2"
)

type ResponsePresenter interface {
	PresentError(c *fiber.Ctx, resp *ErrorResponse) error
	PresentExchangeRate(c *fiber.Ctx, resp *RateResponse) error
	PresentSendRate(c *fiber.Ctx, resp *SendRateResponse) error
}

type SubscriptionUsecase interface {
	Subscribe(recipient *models.Recipient) error
}

type CryptoMailingUsecase interface {
	SendCurrencyRate() ([]string, error)
}

type CryptoExchangerUsecase interface {
	GetCurrentExchangePrice(pair *models.CurrencyPair) (float64, error)
}

type Usecases struct {
	Subscription    SubscriptionUsecase
	CryptoMailing   CryptoMailingUsecase
	CryptoExchanger CryptoExchangerUsecase
}
