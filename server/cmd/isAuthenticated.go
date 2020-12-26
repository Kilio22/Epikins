package main

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/internal/services/loginService"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func isAuthenticated(appData *internal.AppData, c *fiber.Ctx) error {
	accessToken := c.Get("Authorization")
	email, err := loginService.LoginService(appData, accessToken)
	if err != nil {
		return controllers.SendMessage(c, err.Error(), http.StatusUnauthorized)
	}
	c.Request().Header.Set("email", email)
	return c.Next()
}
