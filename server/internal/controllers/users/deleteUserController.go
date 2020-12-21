package users

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/internal/services/users/deleteUserService"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func DeleteUserController(appData *internal.AppData, c *fiber.Ctx) error {
	username := c.Params("username")
	myError := deleteUserService.DeleteUserService(username, appData)
	if myError.Err != nil {
		return controllers.SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.SendStatus(http.StatusNoContent)
}
