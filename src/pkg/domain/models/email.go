package models

type EmailAddress struct {
	Address string `json:"email" validate:"required,email"`
}

type EmailMessage struct {
	Subject string
	Body    string
}
