package users

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/internal/controllers/utils"
	"epikins-api/internal/services/users/addUserService"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AddUserController(appData *internal.AppData, c *fiber.Ctx) error {
	newUser, err := utils.GetUserFromRequest(c)
	if err != nil {
		return controllers.SendMessage(c, err.Error(), http.StatusBadRequest)
	}

	myError := addUserService.AddUserService(newUser, appData)
	if myError.Err != nil {
		return controllers.SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.SendStatus(http.StatusCreated)
}
