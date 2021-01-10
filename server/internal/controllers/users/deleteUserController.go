package users

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/users/deleteUserService"
	"github.com/gofiber/fiber/v2"
)

func DeleteUserController(appData *internal.AppData, c *fiber.Ctx) error {
	username := c.Params("username")
	myError := deleteUserService.DeleteUserService(username, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.SendStatus(http.StatusNoContent)
}
