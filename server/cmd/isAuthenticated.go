package main

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/loginService"
	"github.com/gofiber/fiber/v2"
)

func isAuthenticated(appData *internal.AppData, c *fiber.Ctx) error {
	accessToken := c.Get("Authorization")
	email, myError := loginService.LoginService(appData, accessToken)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	c.Request().Header.Set("email", email)
	return c.Next()
}
