package exchangers

import (
	"genesis_test_case/src/pkg/domain/models"
	"genesis_test_case/src/pkg/domain/usecase"
	"reflect"
)

type loggingExchanger struct {
	exchanger usecase.ExchangeProvider
	logger    usecase.CryptoLogger
}

func NewLoggingExchanger(exc usecase.ExchangeProvider, log usecase.CryptoLogger) *loggingExchanger {
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
