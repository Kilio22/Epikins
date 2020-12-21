package users

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/internal/services/users/getUsersService"
	"github.com/gofiber/fiber/v2"
)

func GetUsersController(appData *internal.AppData, c *fiber.Ctx) error {
	users, myError := getUsersService.GetUsersService(appData)
	if myError.Err != nil {
		return controllers.SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.JSON(users)
}
