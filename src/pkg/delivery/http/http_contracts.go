package http

import (
	"genesis_test_case/src/pkg/delivery/http/responses"
	"github.com/gofiber/fiber/v2"
)

type ResponsePresenter interface {
	PresentError(c *fiber.Ctx, resp *responses.ErrorResponse) error
	PresentExchangeRate(c *fiber.Ctx, resp *responses.RateResponse) error
	PresentSendRate(c *fiber.Ctx, resp *responses.SendRateResponse) error
}
