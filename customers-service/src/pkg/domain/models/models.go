package models

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Recipient
	TransactionId string
	Status        string
}

type Recipient struct {
	EmailAddress string
}
