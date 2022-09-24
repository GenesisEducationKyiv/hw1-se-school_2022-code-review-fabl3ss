package http

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"msg"`
}

type SendRateResponse struct {
	UnsentEmails []string `json:"unsent"`
}

type RateResponse struct {
	Rate float64 `json:"rate"`
}
