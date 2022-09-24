package models

type Recipient struct {
	Email string `json:"email" validate:"required,email"`
}

type EmailMessage struct {
	Subject string
	Body    string
}
