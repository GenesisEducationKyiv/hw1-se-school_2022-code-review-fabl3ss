package domain

import "google.golang.org/api/gmail/v1"

type Recipient struct {
	Email string `json:"email" validate:"required,email"`
}

type MailingUsecase interface {
	SendCurrencyRate() ([]string, error)
	Subscribe(recipient *Recipient) error
}

type MailingRepository interface {
	GetSubscribed() ([]string, error)
	GetMessageBody(bannerURL string) (string, error)
	SendMessage(message *gmail.Message, body string) error
	InsertNewEmail(emails []string, toInsert string) error
	SendToSubscribed(messageBody string) ([]string, error)
}
