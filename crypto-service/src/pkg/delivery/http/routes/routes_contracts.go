package routes

import "github.com/gofiber/fiber/v2"

type MailingHandler interface {
	SendRate(c *fiber.Ctx) error
	Subscribe(c *fiber.Ctx) error
}

type RateHandler interface {
	GetCurrencyRate(c *fiber.Ctx) error
}
