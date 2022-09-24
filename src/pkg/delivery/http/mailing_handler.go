package http

import (
	"errors"
	"genesis_test_case/src/pkg/delivery/http/middleware"
	"genesis_test_case/src/pkg/domain/models"
	myerr "genesis_test_case/src/pkg/types/errors"
	"github.com/gofiber/fiber/v2"
)

type CryptoMailingUsecases struct {
	Exchange     CryptoExchangerUsecase
	Mailing      CryptoMailingUsecase
	Subscription SubscriptionUsecase
}

type MailingHandler struct {
	usecases  *CryptoMailingUsecases
	presenter ResponsePresenter
}

func NewMailingHandler(usecases *CryptoMailingUsecases, presenter ResponsePresenter) *MailingHandler {
	return &MailingHandler{
		usecases:  usecases,
		presenter: presenter,
	}
}

func (m *MailingHandler) SendRate(c *fiber.Ctx) error {
	unsent, err := m.usecases.Mailing.SendCurrencyRate()
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	if len(unsent) > 0 {
		return m.presenter.PresentSendRate(c,
			&SendRateResponse{
				UnsentEmails: unsent,
			},
		)
	}

	return c.SendStatus(fiber.StatusOK)
}

func (m *MailingHandler) Subscribe(c *fiber.Ctx) error {
	recipient := new(models.Recipient)

	errMsg, err := middleware.ParseAndValidate(c, recipient)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errMsg)
	}

	err = m.usecases.Subscription.Subscribe(recipient)
	if err != nil {
		if errors.Is(err, myerr.ErrAlreadyExists) {
			return c.SendStatus(fiber.StatusConflict)
		}
		return m.presenter.PresentError(c,
			&ErrorResponse{
				Error:   true,
				Message: err.Error(),
			},
		)
	}

	return c.SendStatus(fiber.StatusOK)
}
