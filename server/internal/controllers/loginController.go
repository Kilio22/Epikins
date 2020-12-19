package controllers

import (
	"epikins-api/internal"
	"epikins-api/internal/services/loginService"
	"net/http"

	"github.com/gofiber/fiber"
)

func LoginController(appData *internal.AppData, c *fiber.Ctx) {
	accessToken := c.Get("Authorization")
	if _, err := loginService.LoginService(appData.AppId, accessToken); err != nil {
		sendMessage(c, err.Error(), http.StatusUnauthorized)
		return
	}
	c.SendStatus(http.StatusCreated)
}
