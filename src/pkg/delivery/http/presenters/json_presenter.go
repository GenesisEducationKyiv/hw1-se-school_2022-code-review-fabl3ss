package presenters

import (
	"genesis_test_case/src/pkg/delivery/http"
	"github.com/gofiber/fiber/v2"
)

type jsonPresenter struct{}

func NewPresenterJSON() http.ResponsePresenter {
	return &jsonPresenter{}
}

func (j *jsonPresenter) PresentError(c *fiber.Ctx, resp *http.ErrorResponse) error {
	return c.JSON(
		&fiber.Map{
			"error": resp.Error,
			"msg":   resp.Message,
		},
	)
}

func (j *jsonPresenter) PresentExchangeRate(c *fiber.Ctx, resp *http.RateResponse) error {
	return c.JSON(resp.Rate)
}

func (j *jsonPresenter) PresentSendRate(c *fiber.Ctx, resp *http.SendRateResponse) error {
	return c.JSON(
		&fiber.Map{
			"unsent": resp.UnsentEmails,
		},
	)
}
