package jenkinsCredentials

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/jenkinsCredentials/getCredentialsService"
	"github.com/gofiber/fiber/v2"
)

func GetJenkinsCredentialsController(appData *internal.AppData, c *fiber.Ctx) error {
	usernameList, myError := getCredentialsService.GetCredentialsService(appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.JSON(usernameList)
}
