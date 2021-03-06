package jenkinsCredentials

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/jenkinsCredentials/deleteCredentialsService"
	"github.com/gofiber/fiber/v2"
)

func DeleteJenkinsCredentialController(appData *internal.AppData, c *fiber.Ctx) error {
	username := c.Params("username")
	myError := deleteCredentialsService.DeleteCredentialsService(username, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.SendStatus(http.StatusNoContent)
}
