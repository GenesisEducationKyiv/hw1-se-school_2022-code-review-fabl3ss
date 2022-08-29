package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var myValidator = validator.New()

func GetValidator() *validator.Validate {
	return myValidator
}

func ValidatorErrors(err error) map[string]string {
	fields := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		fields[err.Field()] = err.Error()
	}

	return fields
}

func ParseAndValidate(c *fiber.Ctx, payload interface{}) (*fiber.Map, error) {
	if err := c.BodyParser(payload); err != nil {
		return &fiber.Map{
			"error": true,
			"msg":   err.Error(),
		}, err
	}
	if err := myValidator.Struct(payload); err != nil {
		return &fiber.Map{
			"error": true,
			"msg":   ValidatorErrors(err),
		}, err
	}

	return nil, nil
}
