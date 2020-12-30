package controllers

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers/util"
	"epikins-api/internal/services/projectJobsService"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func ProjectJobsController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	projectName := c.Params("project")
	userLogs, err := util.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if err != nil {
		return SendMessage(c, "cannot start builds: "+err.Error(), http.StatusInternalServerError)
	}
	groupsData, myError := projectJobsService.ProjectJobsService(projectName, userLogs, appData)
	if myError.Err != nil {
		return SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.JSON(groupsData)
}
