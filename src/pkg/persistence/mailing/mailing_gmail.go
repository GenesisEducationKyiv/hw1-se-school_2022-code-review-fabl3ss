package mailing

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"

	"genesis_test_case/src/config"
	"genesis_test_case/src/pkg/domain"
	"genesis_test_case/src/pkg/usecase"
	"genesis_test_case/src/platform/gmail_api"

	"github.com/pkg/errors"
	"google.golang.org/api/gmail/v1"
)

type mailingGmailRepository struct {
	gmailService *gmail.Service
}

func NewGmailRepository(srv *gmail.Service) usecase.MailingRepository {
	return &mailingGmailRepository{
		gmailService: srv,
	}
}

func (m *mailingGmailRepository) MultipleSending(msg *domain.EmailMessage, recipients []string) ([]string, error) {
	var (
		unsent  []string
		message gmail.Message
	)

	messageBody := fmt.Sprintf(
		"Subject: %s\r\n"+
			"MIME-Version: 1.0\r\n"+
			"Content-Type: text/html; charset=UTF-8\r\n"+
			"Content-Transfer-Encoding: base64\r\n\r\n"+
			"%s",
		msg.Subject, msg.Body,
	)

	for _, recipient := range recipients {
		err := m.sendMessage(&message, fmt.Sprintf("To: %s\r\n", recipient)+messageBody)
		if err != nil {
			log.Printf("Sending error!: %v\n", err)
			unsent = append(unsent, recipient)
		} else {
			log.Printf("Message sent! to %s\n", recipient)
		}
	}

	return unsent, nil
}

func (m *mailingGmailRepository) sendMessage(message *gmail.Message, body string) error {
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

func (m *mailingGmailRepository) updateOauthToken() error {
	service, err := gmail_api.UpdateService(
		os.Getenv(config.EnvGmailCredentialsPath),
		os.Getenv(config.EnvGmailTokenPath),
	)
	if err != nil {
		return err
	}

	m.gmailService = service

	return nil
}
