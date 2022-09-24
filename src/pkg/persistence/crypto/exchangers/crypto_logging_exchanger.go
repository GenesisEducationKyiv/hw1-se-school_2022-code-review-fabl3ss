package exchangers

import (
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/models"
	"reflect"
)

type loggingExchanger struct {
	exchanger application.ExchangeProvider
	logger    application.CryptoLogger
}

func NewLoggingExchanger(exc application.ExchangeProvider, log application.CryptoLogger) *loggingExchanger {
	return &loggingExchanger{
		exchanger: exc,
		logger:    log,
	}
}

func (l *loggingExchanger) GetCurrencyRate(pair *models.CurrencyPair) (*models.CurrencyRate, error) {
	rate, err := l.exchanger.GetCurrencyRate(pair)
	if err != nil {
		return nil, err
	}

	l.logger.LogExchangeRate(l.getExhangerName(), rate)

	return rate, nil
}

func (l *loggingExchanger) getExhangerName() string {
	return reflect.TypeOf(l.exchanger).Elem().Name()
}
