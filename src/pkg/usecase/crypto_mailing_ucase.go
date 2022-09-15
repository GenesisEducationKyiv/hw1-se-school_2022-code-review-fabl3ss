package usecase

import (
	"genesis_test_case/src/pkg/delivery/http"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/utils"
)

type CryptoMailingRepositories struct {
	Repositories
}

type cryptoMailingUsecase struct {
	pair             *domain.CurrencyPair
	repos            *CryptoMailingRepositories
	templatePathHTML string
}

func NewCryptoMailingUsecase(
	htmlPath string,
	pair *domain.CurrencyPair,
	repos *CryptoMailingRepositories,
) http.CryptoMailingUsecase {
	return &cryptoMailingUsecase{
		pair:             pair,
		repos:            repos,
		templatePathHTML: htmlPath,
	}
}

func (c *cryptoMailingUsecase) SendCurrencyRate() ([]string, error) {
	bannerURL, err := c.getMailingBannerUrl()
	if err != nil {
		return nil, err
	}
	messageBody, err := c.buildMessage(bannerURL)
	if err != nil {
		return nil, err
	}

	return c.sendToSubscribed(messageBody)
}

func (c *cryptoMailingUsecase) sendToSubscribed(message *domain.EmailMessage) ([]string, error) {
	recipients, err := c.repos.Storage.GetAllEmails()
	if err != nil {
		return nil, err
	}
	unsent, err := c.repos.Mailer.MultipleSending(message, recipients)
	return unsent, err
}

func (c *cryptoMailingUsecase) getMailingBannerUrl() (string, error) {
	chart, err := c.repos.Chart.GetWeekAverageChart(c.pair)
	if err != nil {
		return "", err
	}
	rate, err := c.repos.Exchanger.GetCurrencyRate(c.pair)
	if err != nil {
		return "", err
	}

	return c.repos.Banner.GetCryptoBannerUrl(chart, rate)
}

func (c *cryptoMailingUsecase) buildMessage(bannerURL string) (*domain.EmailMessage, error) {
	v := struct {
		Chart string
	}{Chart: bannerURL}

	htmlContent, err := utils.ParseHtmlTemplate(c.templatePathHTML, &v)
	if err != nil {
		return nil, err
	}

	return &domain.EmailMessage{
		Subject: "Crypto Newsletter",
		Body:    htmlContent.String(),
	}, nil
}
