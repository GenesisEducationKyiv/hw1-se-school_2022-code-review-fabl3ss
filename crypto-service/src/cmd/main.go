package main

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/delivery/http/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New(config.FiberConfig())

	err := routes.InitRoutes(app)
	if err != nil {
		panic("Unable to initialize routes")
	}
	http.StartServerWithGracefulShutdown(app)
}
