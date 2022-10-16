package config

import (
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

const (
	EnvServerUrl         string = "SERVER_URL"
	EnvServerReadTimeout string = "SERVER_READ_TIMEOUT"
)

func FiberConfig() fiber.Config {
	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv(EnvServerReadTimeout))

	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	}
}
