package usecase

import "genesis_test_case/src/pkg/domain"

type MailingRepository interface {
	MultipleSending(message *domain.EmailMessage, adresses []string) ([]string, error)
}

type EmailStorage interface {
	GetAllEmails() ([]string, error)
	AddEmail(email string) error
}

type ExchangeProvider interface {
	GetCurrencyRate(pair *domain.CurrencyPair) (*domain.CurrencyRate, error)
}

type ChartProvider interface {
	GetWeekAverageChart(pair *domain.CurrencyPair) ([]float64, error)
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
	GetCryptoBannerUrl(chart []float64, rate *domain.CurrencyRate) (string, error)
}

type Cache interface {
	GetCache(key string) ([]byte, error)
	SetCache(key string, value interface{}) error
}

type CryptoCache interface {
	SetCurrencyCache(key string, rate *domain.CurrencyRate) error
	GetCurrencyCache(key string) (*domain.CurrencyRate, error)
}

type CryptoLogger interface {
	LogExchangeRate(provider string, rate *domain.CurrencyRate)
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
