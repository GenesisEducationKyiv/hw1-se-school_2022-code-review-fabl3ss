package crypto

import (
	"genesis_test_case/src/loggers"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/usecase"
)

type cryptoLogger struct {
	logger loggers.Logger
}

func NewCryptoLogger(logger loggers.Logger) usecase.CryptoLogger {
	return &cryptoLogger{
		logger: logger,
	}
}

func (c *cryptoLogger) LogExchangeRate(provider string, rate *domain.CurrencyRate) {
	c.logger.Infow(
		"recieved rate",
		"provider", provider,
		"price", rate.Price,
		"base", rate.GetBaseCurrency(),
		"quote", rate.GetQuoteCurrency(),
	)
}
