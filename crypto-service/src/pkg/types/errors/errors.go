package errors

import (
	"github.com/pkg/errors"
)

var (
	ErrAlreadyExists       = errors.New("record already exists")
	ErrFailedParseHttpBody = errors.New("failed parse HTTP body")
	ErrValidationFailed    = errors.New("validation falied")
	ErrInvalidInput        = errors.New("invalid input")
	ErrNoDataProvided      = errors.New("no data provided")
	ErrNoSubscribers       = errors.New("no subscribers in storage")
	ErrNoCache             = errors.New("cache is empty")
)
