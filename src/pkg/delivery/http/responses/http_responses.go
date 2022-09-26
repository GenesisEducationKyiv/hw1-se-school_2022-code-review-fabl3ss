package responses

import "genesis_test_case/src/pkg/domain/models"

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"msg"`
}

type SendRateResponse struct {
	UnsentEmails []models.EmailAddress `json:"unsent"`
}

type RateResponse struct {
	Rate float64 `json:"rate"`
}
