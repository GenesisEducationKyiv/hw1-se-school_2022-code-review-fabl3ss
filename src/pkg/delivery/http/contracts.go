package http

import (
	"genesis_test_case/src/pkg/domain"
)

type SubscriptionUsecase interface {
	Subscribe(recipient *domain.Recipient) error
}

type CryptoMailingUsecase interface {
	SendCurrencyRate() ([]string, error)
}

type CryptoExchangerUsecase interface {
	GetCurrentExchangePrice() (float64, error)
}

type Usecases struct {
	Subscription    SubscriptionUsecase
	CryptoMailing   CryptoMailingUsecase
	CryptoExchanger CryptoExchangerUsecase
}
