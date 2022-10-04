package presenters

import (
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/delivery/http/responses"

	"github.com/gofiber/fiber/v2"
)

type jsonPresenter struct{}

func NewPresenterJSON() http.ResponsePresenter {
	return &jsonPresenter{}
}

func (j *jsonPresenter) PresentError(c *fiber.Ctx, resp *responses.ErrorResponse) error {
	return c.JSON(
		&fiber.Map{
			"error": resp.Error,
			"msg":   resp.Message,
		},
	)
}

func (j *jsonPresenter) PresentExchangeRate(c *fiber.Ctx, resp *responses.RateResponse) error {
	return c.JSON(resp.Rate)
}

func (j *jsonPresenter) PresentSendRate(c *fiber.Ctx, resp *responses.SendRateResponse) error {
	return c.JSON(
		&fiber.Map{
			"unsent": resp.UnsentEmails,
		},
	)
}
