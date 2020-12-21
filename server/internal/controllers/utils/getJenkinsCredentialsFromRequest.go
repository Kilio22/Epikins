package utils

import (
	"epikins-api/internal"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetJenkinsCredentialsFromRequest(c *fiber.Ctx) (internal.JenkinsCredentials, error) {
	var credentials internal.JenkinsCredentials

	err := c.BodyParser(&credentials)
	if err != nil {
		return internal.JenkinsCredentials{}, errors.New("wrong body")
	}
	err = validator.New().Struct(credentials)
	if err != nil {
		return internal.JenkinsCredentials{}, errors.New("wrong body")
	}
	return credentials, nil
}
