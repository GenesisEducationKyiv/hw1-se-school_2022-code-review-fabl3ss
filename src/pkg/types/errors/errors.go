package errors

import (
	"github.com/pkg/errors"
)

var ErrAlreadyExists = errors.New("record already exists")
