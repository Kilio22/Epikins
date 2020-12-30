package users

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/internal/controllers/util"
	"epikins-api/internal/services/users/updateUserService"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func UpdateUserController(appData *internal.AppData, c *fiber.Ctx) error {
	user, err := util.GetUserFromRequest(c)
	if err != nil {
		return controllers.SendMessage(c, err.Error(), http.StatusBadRequest)
	}

	myError := updateUserService.UpdateUserService(user, appData)
	if myError.Err != nil {
		return controllers.SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.SendStatus(http.StatusCreated)
}
