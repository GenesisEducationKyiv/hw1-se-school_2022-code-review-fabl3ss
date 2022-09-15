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

type CryptoBannerRepository interface {
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

type Repositories struct {
	Banner    CryptoBannerRepository
	Exchanger ExchangeProvider
	Storage   EmailStorage
	Mailer    MailingRepository
}
