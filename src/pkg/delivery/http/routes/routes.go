package routes

import (
	"github.com/gofiber/fiber/v2"
)

type Handlers struct {
	Mailing MailingHandler
	Rate    RateHandler
}

func InitPublicRoutes(a *fiber.App, handlers *Handlers) {
	a.Get("/rate", handlers.Rate.GetCurrencyRate)
	a.Post("/sendEmails", handlers.Mailing.SendRate)
	a.Post("/subscribe", handlers.Mailing.Subscribe)
}
