package usecase

import "genesis_test_case/src/pkg/domain"

type MailingRepository interface {
	MultipleSending(message *domain.EmailMessage, adresses []string) ([]string, error)
}

type EmailStorage interface {
	GetAllEmails() ([]string, error)
	AddEmail(email string) error
}

type CryptoRepository interface {
	GetCurrencyRate(pair *domain.CurrencyPair) (*domain.CurrencyRate, error)
	GetWeekAverageChart(pair *domain.CurrencyPair) ([]float64, error)
}

type CryptoBannerRepository interface {
	GetCryptoBannerUrl(chart []float64, rate *domain.CurrencyRate) (string, error)
}

type Repositories struct {
	Banner    CryptoBannerRepository
	Exchanger CryptoRepository
	Storage   EmailStorage
	Mailer    MailingRepository
}
