package http

import (
	"errors"
	"genesis_test_case/src/pkg/delivery/http/responses"
	"genesis_test_case/src/pkg/domain"
	myerr "genesis_test_case/src/pkg/types/errors"
	"genesis_test_case/src/pkg/usecase"
	"genesis_test_case/src/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type MailingHandler struct {
	usecases *usecase.Usecases
}

func NewMailingHandler(u *usecase.Usecases) *MailingHandler {
	return &MailingHandler{
		usecases: u,
	}
}

func (m *MailingHandler) SendRate(c *fiber.Ctx) error {
	unsent, err := m.usecases.Mailing.SendCurrencyRate()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if len(unsent) > 0 {
		return c.JSON(
			responses.SendRateResponseHTTP{
				UnsentEmails: unsent,
			},
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (m *MailingHandler) Subscribe(c *fiber.Ctx) error {
	recipient := new(domain.Recipient)

	errMsg, err := utils.ParseAndValidate(c, recipient)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	err = m.usecases.Mailing.Subscribe(recipient)
	if err != nil {
		if errors.Is(err, myerr.ErrAlreadyExists) {
			return c.SendStatus(fiber.StatusConflict)
		}
		return c.Status(fiber.StatusInternalServerError).JSON(
			responses.ErrorResponseHTTP{
				Error:   true,
				Message: err.Error(),
			},
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (m *MailingHandler) GetCurrencyRate(c *fiber.Ctx) error {
	rate, err := m.usecases.Crypto.GetConfigCurrencyRate()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(rate)
}
