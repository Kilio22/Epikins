package users

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/users/getUsersService"
	"github.com/gofiber/fiber/v2"
)

func GetUsersController(appData *internal.AppData, c *fiber.Ctx) error {
	users, myError := getUsersService.GetUsersService(appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.JSON(users)
}
