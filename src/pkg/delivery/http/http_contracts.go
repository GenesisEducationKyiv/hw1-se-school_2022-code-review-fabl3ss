package http

import (
	"genesis_test_case/src/pkg/delivery/http/presentation"
	"genesis_test_case/src/pkg/domain"
	"github.com/gofiber/fiber/v2"
)

type ResponsePresenter interface {
	PresentError(c *fiber.Ctx, resp *presentation.ErrorResponse) error
	PresentExchangeRate(c *fiber.Ctx, resp *presentation.RateResponse) error
	PresentSendRate(c *fiber.Ctx, resp *presentation.SendRateResponse) error
}

type SubscriptionUsecase interface {
	Subscribe(recipient *domain.Recipient) error
}

type CryptoMailingUsecase interface {
	SendCurrencyRate() ([]string, error)
}

type CryptoExchangerUsecase interface {
	GetCurrentExchangePrice(pair *domain.CurrencyPair) (float64, error)
}

type Usecases struct {
	Subscription    SubscriptionUsecase
	CryptoMailing   CryptoMailingUsecase
	CryptoExchanger CryptoExchangerUsecase
}
