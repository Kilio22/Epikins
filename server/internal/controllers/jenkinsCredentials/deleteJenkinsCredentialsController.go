package jenkinsCredentials

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/internal/services/jenkinsCredentials/deleteCredentialsService"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func DeleteJenkinsCredentialController(appData *internal.AppData, c *fiber.Ctx) error {
	username := c.Params("username")
	myError := deleteCredentialsService.DeleteCredentialsService(username, appData)
	if myError.Err != nil {
		return controllers.SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.SendStatus(http.StatusNoContent)
}
