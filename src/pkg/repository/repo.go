package repository

import (
	"genesis_test_case/src/pkg/domain"

	"google.golang.org/api/gmail/v1"
)

type Repositories struct {
	Image   domain.ImageRepository
	Crypto  domain.CryptoRepository
	Mailing domain.MailingRepository
}

func NewRepositories(gmailService *gmail.Service) *Repositories {
	return &Repositories{
		Image:   NewImageRepository(),
		Crypto:  NewCryptoRepository(),
		Mailing: NewMailingRepository(gmailService),
	}
}
