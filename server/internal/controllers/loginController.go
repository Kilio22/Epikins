package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}
