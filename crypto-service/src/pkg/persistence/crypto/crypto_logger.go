package crypto

import (
	"encoding/json"
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/logger"
	"genesis_test_case/src/pkg/domain/models"
	"log"
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
	marshal, err := json.Marshal(map[string]any{
		"provider": provider,
		"price":    rate.Price,
		"base":     rate.GetBaseCurrency(),
		"quote":    rate.GetQuoteCurrency(),
	})
	if err != nil {
		log.Println("unable to marshal exchange log")
	}

	c.logger.Info(string(marshal))
}
