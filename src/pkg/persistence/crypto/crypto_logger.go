package crypto

import (
	"genesis_test_case/src/loggers"
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/models"
)

type cryptoLogger struct {
	logger loggers.Logger
}

func NewCryptoLogger(logger loggers.Logger) application.CryptoLogger {
	return &cryptoLogger{
		logger: logger,
	}
}

func (c *cryptoLogger) LogExchangeRate(provider string, rate *models.CurrencyRate) {
	c.logger.Infow(
		"recieved rate",
		"provider", provider,
		"price", rate.Price,
		"base", rate.GetBaseCurrency(),
		"quote", rate.GetQuoteCurrency(),
	)
}
