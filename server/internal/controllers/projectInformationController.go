package controllers

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/projectInformationService"
	"epikins-api/internal/services/projectsService"
	"epikins-api/internal/services/util"
	"github.com/gofiber/fiber/v2"
)

func ProjectInformationController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	projectName := c.Params("project")

	userLogs, err := controllerUtil.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(projectsService.ProjectsError, err, http.StatusInternalServerError), c)
	}
	projectList, myError := projectInformationService.ProjectInformationService(projectName, userLogs, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.JSON(projectList)
}
