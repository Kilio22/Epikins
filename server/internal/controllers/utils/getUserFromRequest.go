package utils

import (
	"epikins-api/internal"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetUserFromRequest(c *fiber.Ctx) (internal.User, error) {
	var user internal.User

	err := c.BodyParser(&user)
	if err != nil {
		return internal.User{}, errors.New("wrong body")
	}
	err = validator.New().Struct(user)
	if err != nil {
		return internal.User{}, errors.New("wrong body")
	}
	return user, nil
}
