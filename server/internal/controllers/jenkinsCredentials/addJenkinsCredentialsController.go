package jenkinsCredentials

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/internal/controllers/util"
	"epikins-api/internal/services/jenkinsCredentials/addCredentialsService"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func AddJenkinsCredentialController(appData *internal.AppData, c *fiber.Ctx) error {
	newCredentials, err := util.GetJenkinsCredentialsFromRequest(c)
	if err != nil {
		return controllers.SendMessage(c, err.Error(), http.StatusBadRequest)
	}

	myError := addCredentialsService.AddCredentialsService(newCredentials, appData)
	if myError.Err != nil {
		return controllers.SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.SendStatus(http.StatusCreated)
}
