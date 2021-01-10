package controllerUtil

import (
	"errors"
	"strings"

	"epikins-api/pkg/libJenkins"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetJenkinsCredentialsFromRequest(c *fiber.Ctx) (libJenkins.JenkinsCredentials, error) {
	var credentials libJenkins.JenkinsCredentials

	err := c.BodyParser(&credentials)
	if err != nil {
		return libJenkins.JenkinsCredentials{}, errors.New("wrong body")
	}
	credentials.Login = strings.TrimSpace(credentials.Login)
	credentials.ApiKey = strings.TrimSpace(credentials.ApiKey)
	err = validator.New().Struct(credentials)
	if err != nil {
		return libJenkins.JenkinsCredentials{}, errors.New("wrong body")
	}
	return credentials, nil
}
