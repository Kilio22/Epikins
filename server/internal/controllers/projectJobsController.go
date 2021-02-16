package controllers

import (
	"net/http"
	"strings"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/projectJobsService"
	"epikins-api/internal/services/util"
	"github.com/gofiber/fiber/v2"
)

func ProjectJobsController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	projectName := c.Params("project")
	module := c.Params("module")
	city := strings.ToUpper(c.Params("city"))

	userLogs, err := controllerUtil.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(projectJobsService.ProjectJobsError, err, http.StatusInternalServerError), c)
	}
	workgroupsData, myError := projectJobsService.ProjectJobsService(projectName, module, city, userLogs, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.JSON(workgroupsData)
}
