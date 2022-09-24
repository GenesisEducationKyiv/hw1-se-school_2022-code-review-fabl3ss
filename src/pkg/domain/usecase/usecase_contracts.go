package usecase

import (
	"genesis_test_case/src/pkg/domain/models"
)

type MailingRepository interface {
	MultipleSending(message *models.EmailMessage, adresses []string) ([]string, error)
}

type EmailStorage interface {
	GetAllEmails() ([]string, error)
	AddEmail(email string) error
}

type ExchangeProvider interface {
	GetCurrencyRate(pair *models.CurrencyPair) (*models.CurrencyRate, error)
}

type ChartProvider interface {
	GetWeekAverageChart(pair *models.CurrencyPair) ([]float64, error)
}

type ExchangeProviderNode interface {
	ExchangeProvider
	SetNext(exchanger ExchangeProviderNode)
}

type ExchangersChain interface {
	RegisterExchanger(name string, exchanger, next ExchangeProviderNode) error
	GetExchanger(name string) ExchangeProvider
}

type CryptoBannerProvider interface {
	GetCryptoBannerUrl(chart []float64, rate *models.CurrencyRate) (string, error)
}

type Cache interface {
	GetCache(key string) ([]byte, error)
	SetCache(key string, value interface{}) error
}

type CryptoCache interface {
	SetCurrencyCache(key string, rate *models.CurrencyRate) error
	GetCurrencyCache(key string) (*models.CurrencyRate, error)
}

type CryptoLogger interface {
	LogExchangeRate(provider string, rate *models.CurrencyRate)
}

type Repositories struct {
	Banner    CryptoBannerProvider
	Exchanger ExchangeProvider
	Chart     ChartProvider
	Storage   EmailStorage
	Mailer    MailingRepository
}

type CryptoMailingRepositories struct {
	Repositories
}
