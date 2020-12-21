package jenkinsCredentials

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/internal/controllers/utils"
	"epikins-api/internal/services/jenkinsCredentials/updateCredentialsService"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func UpdateJenkinsCredentialsController(appData *internal.AppData, c *fiber.Ctx) error {
	newCredentials, err := utils.GetJenkinsCredentialsFromRequest(c)
	if err != nil {
		return controllers.SendMessage(c, err.Error(), http.StatusBadRequest)
	}

	myError := updateCredentialsService.UpdateCredentialsService(newCredentials, appData)
	if myError.Err != nil {
		return controllers.SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.SendStatus(http.StatusCreated)
}
