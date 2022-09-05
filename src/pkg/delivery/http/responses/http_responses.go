package responses

type ErrorResponseHTTP struct {
	Error   bool   `json:"error"`
	Message string `json:"msg"`
}

type SendRateResponseHTTP struct {
	UnsentEmails []string `json:"unsent"`
}
