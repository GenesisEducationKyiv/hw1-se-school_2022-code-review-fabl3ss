package errors

import (
	"github.com/pkg/errors"
)

var ErrAlreadyExists = errors.New("record already exists")
var ErrFailedParseHttpBody = errors.New("failed parse HTTP body")
var ErrValidationFailed = errors.New("validation falied")
var ErrInvalidInput = errors.New("invalid input")
var ErrNoDataProvided = errors.New("no data provided")
var ErrNoSubscribers = errors.New("no subscribers in storage")
