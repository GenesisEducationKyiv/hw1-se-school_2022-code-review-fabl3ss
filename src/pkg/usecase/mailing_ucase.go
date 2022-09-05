package usecase

import (
	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/repository"

	myerr "genesis_test_case/src/pkg/types/errors"
)

type mailingUsecase struct {
	repos *repository.Repositories
}

func NewMailingUsecase(r *repository.Repositories) domain.MailingUsecase {
	return &mailingUsecase{
		repos: r,
	}
}

func (m *mailingUsecase) Subscribe(recipient *domain.Recipient) error {
	if recipient == nil {
		return myerr.ErrNoDataProvided
	}
	subscribed, err := m.repos.Mailing.GetSubscribed()
	if err != nil {
		return err
	}
	err = m.repos.Mailing.InsertNewEmail(subscribed, recipient.Email)
	if err != nil {
		return err
	}

	return nil
}

func (m *mailingUsecase) SendCurrencyRate() ([]string, error) {
	bannerURL, err := m.getMailingBannerUrl()
	if err != nil {
		return nil, err
	}
	messageBody, err := m.repos.Mailing.GetMessageBody(bannerURL)
	if err != nil {
		return nil, err
	}
	unsent, err := m.repos.Mailing.SendToSubscribed(messageBody)

	return unsent, err
}

func (m *mailingUsecase) getMailingBannerUrl() (string, error) {
	cfg := config.Get()
	chart, err := m.repos.Crypto.GetWeekChart()
	if err != nil {
		return "", err
	}
	rate, err := m.repos.Crypto.GetCurrencyRate(cfg.BaseCurrency, cfg.QuoteCurrency)
	if err != nil {
		return "", err
	}

	return m.repos.Image.GetCryptoBannerUrl(chart, rate)
}
