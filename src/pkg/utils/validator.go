package utils

import (
	"bytes"
	"fmt"
	"genesis_test_case/src/pkg/delivery/http/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var myValidator = validator.New()

func GetValidator() *validator.Validate {
	return myValidator
}

func validatorErrors(err error) string {
	errorMessage := new(bytes.Buffer)

	for _, err := range err.(validator.ValidationErrors) {
		fmt.Fprintf(errorMessage, "field: %s; error: %s\n", err.Field(), err.Error())
	}

	return errorMessage.String()
}

func ParseAndValidate(c *fiber.Ctx, payload interface{}) (*responses.ErrorResponseHTTP, error) {
	if err := c.BodyParser(payload); err != nil {
		return &responses.ErrorResponseHTTP{
			Error:   true,
			Message: err.Error(),
		}, err
	}
	if err := myValidator.Struct(payload); err != nil {
		return &responses.ErrorResponseHTTP{
			Error:   true,
			Message: validatorErrors(err),
		}, err
	}

	return nil, nil
}
