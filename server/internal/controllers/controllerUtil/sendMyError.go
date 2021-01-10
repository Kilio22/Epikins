package controllerUtil

import (
	"epikins-api/internal"
	"github.com/gofiber/fiber/v2"
)

func SendMyError(myError internal.MyError, c *fiber.Ctx) error {
	c.Status(myError.Status)
	return c.JSON(myError)
}
