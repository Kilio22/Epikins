package users

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/users/addUserService"
	"epikins-api/internal/services/util"
	"github.com/gofiber/fiber/v2"
)

func AddUserController(appData *internal.AppData, c *fiber.Ctx) error {
	newUser, err := controllerUtil.GetUserFromRequest(c)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(addUserService.AddUserError, err, http.StatusBadRequest), c)
	}

	myError := addUserService.AddUserService(newUser, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.SendStatus(http.StatusCreated)
}
