package users

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/users/updateUserService"
	"epikins-api/internal/services/util"
	"github.com/gofiber/fiber/v2"
)

func UpdateUserController(appData *internal.AppData, c *fiber.Ctx) error {
	user, err := controllerUtil.GetUserFromRequest(c)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(updateUserService.UpdateUserError, err, http.StatusBadRequest), c)
	}

	myError := updateUserService.UpdateUserService(user, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.SendStatus(http.StatusCreated)
}
