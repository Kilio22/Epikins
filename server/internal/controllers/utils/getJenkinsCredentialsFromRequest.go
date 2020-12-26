package utils

import (
	"epikins-api/pkg/libJenkins"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetJenkinsCredentialsFromRequest(c *fiber.Ctx) (libJenkins.JenkinsCredentials, error) {
	var credentials libJenkins.JenkinsCredentials

	err := c.BodyParser(&credentials)
	if err != nil {
		return libJenkins.JenkinsCredentials{}, errors.New("wrong body")
	}
	err = validator.New().Struct(credentials)
	if err != nil {
		return libJenkins.JenkinsCredentials{}, errors.New("wrong body")
	}
	return credentials, nil
}
