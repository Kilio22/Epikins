package jenkinsCredentials

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers"
	"epikins-api/internal/services/jenkinsCredentials/getCredentialsService"
	"github.com/gofiber/fiber/v2"
)

func GetJenkinsCredentialsController(appData *internal.AppData, c *fiber.Ctx) error {
	usernameList, myError := getCredentialsService.GetCredentialsService(appData)
	if myError.Err != nil {
		return controllers.SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.JSON(usernameList)
}
