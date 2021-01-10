package controllerUtil

import (
	"errors"
	"strings"

	"epikins-api/config"
	"epikins-api/internal"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func trimValues(user *internal.User) {
	user.JenkinsLogin = strings.TrimSpace(user.JenkinsLogin)
	user.Email = strings.TrimSpace(user.Email)
	for idx, role := range user.Roles {
		user.Roles[idx] = internal.Role(strings.TrimSpace(string(role)))
	}
}

func areRolesValid(roles []internal.Role) bool {
	var isValid = false

	for _, role := range roles {
		for _, value := range config.Roles {
			if role == value {
				isValid = true
				break
			}
		}
		if !isValid {
			return false
		}
		isValid = false
	}
	return true
}

func GetUserFromRequest(c *fiber.Ctx) (internal.User, error) {
	var user internal.User

	err := c.BodyParser(&user)
	if err != nil {
		return internal.User{}, errors.New("wrong body")
	}
	trimValues(&user)
	err = validator.New().Struct(user)
	if err != nil || !areRolesValid(user.Roles) {
		return internal.User{}, errors.New("wrong body")
	}
	return user, nil
}
