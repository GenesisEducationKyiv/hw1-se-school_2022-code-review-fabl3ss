package usecase

import (
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/repository"
)

type Usecases struct {
	Mailing domain.MailingUsecase
	Crypto  domain.CryptoUsecase
}

func NewUsecases(repos *repository.Repositories) *Usecases {
	return &Usecases{
		Mailing: NewMailingUsecase(repos),
		Crypto:  NewCryptoUsecase(repos),
	}
}
