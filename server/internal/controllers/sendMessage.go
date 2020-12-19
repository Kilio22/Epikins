package controllers

import "github.com/gofiber/fiber"

func sendMessage(c *fiber.Ctx, message string, status int) {
	c.Status(status)
	c.Send(message)
}
