package routes

import (
	"genesis_test_case/src/pkg/delivery/http"

	"github.com/gofiber/fiber/v2"
)

func InitPublicRoutes(a *fiber.App, handler *http.MailingHandler) {
	a.Get("/rate", handler.GetCurrencyRate)
	a.Post("/sendEmails", handler.SendRate)
	a.Post("/subscribe", handler.Subscribe)
}
