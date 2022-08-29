package usecase

import (
	domain2 "genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/repository"
)

type Usecases struct {
	Mailing domain2.MailingUsecase
	Crypto  domain2.CryptoUsecase
}

func NewUsecases(repos *repository.Repositories) *Usecases {
	return &Usecases{
		Mailing: NewMailingUsecase(repos),
		Crypto:  NewCryptoUsecase(repos),
	}
}
