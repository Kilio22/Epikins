package controllers

import "github.com/gofiber/fiber/v2"

func SendMessage(c *fiber.Ctx, message string, status int) error {
	c.Status(status)
	return c.SendString(message)
}
