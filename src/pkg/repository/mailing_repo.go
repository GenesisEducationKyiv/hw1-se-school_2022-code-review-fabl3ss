package repository

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"html/template"
	"log"
	"os"

	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/utils"
	"genesis_test_case/src/platform/gmail_api"

	"github.com/pkg/errors"
	"google.golang.org/api/gmail/v1"
)

type mailingRepository struct {
	gmailService *gmail.Service
}

func NewMailingRepository(srv *gmail.Service) domain.MailingRepository {
	return &mailingRepository{
		gmailService: srv,
	}
}

func (m *mailingRepository) updateOauthToken() error {
	service, err := gmail_api.UpdateService(
		os.Getenv(config.GmailCredentialsPath),
		os.Getenv(config.GmailTokenPath),
	)
	if err != nil {
		return err
	}

	m.gmailService = service

	return nil
}

func (m *mailingRepository) GetSubscribed() ([]string, error) {
	cfg := config.Get()

	subscribed, err := utils.ReadAllFromCsv(cfg.StorageFile)
	if err != nil {
		return nil, err
	}

	return subscribed, nil
}

func (m *mailingRepository) SendMessage(message *gmail.Message, body string) error {
	if err := m.updateOauthToken(); err != nil {
		return err
	}

	message.Raw = base64.URLEncoding.EncodeToString([]byte(body))
	_, err := m.gmailService.Users.Messages.Send("me", message).Do()
	if err != nil {
		return errors.Wrap(err, "sending failed")
	}

	return nil
}

func (m *mailingRepository) SendToSubscribed(messageBody string) ([]string, error) {
	var unsent []string
	var message gmail.Message
	subscribed, err := m.GetSubscribed()
	if err != nil {
		return nil, err
	}

	for _, email := range subscribed {
		err = m.SendMessage(&message, fmt.Sprintf("To: %s\r\n", email)+messageBody)
		if err != nil {
			log.Printf("Sending error!: %v\n", err)
			unsent = append(unsent, email)
		} else {
			log.Printf("Message sent! to %s\n", email)
		}
	}

	return unsent, nil
}

func (m *mailingRepository) InsertNewEmail(emails []string, toInsert string) error {
	cfg := config.Get()

	emails, err := utils.InsertToSorted(emails, toInsert)
	if err != nil {
		return err
	}
	err = utils.WriteToCsv(cfg.StorageFile, emails)
	if err != nil {
		return err
	}

	return nil
}

func (m *mailingRepository) GetMessageBody(bannerURL string) (string, error) {
	var htmlContent bytes.Buffer
	v := struct {
		Chart string
	}{Chart: bannerURL}
	t, _ := template.ParseFiles(
		os.Getenv(config.CryptoHtmlMessagePath),
	)

	err := t.Execute(&htmlContent, v)
	if err != nil {
		return "", err
	}
	messageBody := "Subject: Crypto Newsletter\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"Content-Transfer-Encoding: base64\r\n\r\n" +
		htmlContent.String()

	return messageBody, nil
}
