package http

type ResponseError struct {
	Error   bool   `json:"error"`
	Message string `json:"msg"`
}
