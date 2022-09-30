package crypto

import (
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/logger"
	"genesis_test_case/src/pkg/domain/models"
)

type cryptoLogger struct {
	logger logger.Logger
}

func NewCryptoLogger(logger logger.Logger) application.CryptoLogger {
	return &cryptoLogger{
		logger: logger,
	}
}

func (c *cryptoLogger) LogExchangeRate(provider string, rate *models.CurrencyRate) {
	c.logger.Infow(
		"received rate",
		"provider", provider,
		"price", rate.Price,
		"base", rate.GetBaseCurrency(),
		"quote", rate.GetQuoteCurrency(),
	)
}
