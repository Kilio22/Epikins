package jenkinsCredentials

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/jenkinsCredentials/addCredentialsService"
	"epikins-api/internal/services/util"
	"github.com/gofiber/fiber/v2"
)

func AddJenkinsCredentialController(appData *internal.AppData, c *fiber.Ctx) error {
	newCredentials, err := controllerUtil.GetJenkinsCredentialsFromRequest(c)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(addCredentialsService.AddCredentialsError, err, http.StatusBadRequest), c)
	}

	myError := addCredentialsService.AddCredentialsService(newCredentials, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.SendStatus(http.StatusCreated)
}
