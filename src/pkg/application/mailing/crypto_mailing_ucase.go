package usecase

import (
	"genesis_test_case/src/pkg/application"
	"genesis_test_case/src/pkg/domain/models"
	"genesis_test_case/src/pkg/domain/usecases"
	"genesis_test_case/src/pkg/utils"
)

type cryptoMailingUsecase struct {
	pair             *models.CurrencyPair
	repos            *application.CryptoMailingRepositories
	templatePathHTML string
}

func NewCryptoMailingUsecase(
	htmlPath string,
	pair *models.CurrencyPair,
	repos *application.CryptoMailingRepositories,
) usecases.CryptoMailingUsecase {
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

func (c *cryptoMailingUsecase) sendToSubscribed(message *models.EmailMessage) ([]string, error) {
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

func (c *cryptoMailingUsecase) buildMessage(bannerURL string) (*models.EmailMessage, error) {
	v := struct {
		Chart string
	}{Chart: bannerURL}

	htmlContent, err := utils.ParseHtmlTemplate(c.templatePathHTML, &v)
	if err != nil {
		return nil, err
	}

	return &models.EmailMessage{
		Subject: "Crypto Newsletter",
		Body:    htmlContent.String(),
	}, nil
}
