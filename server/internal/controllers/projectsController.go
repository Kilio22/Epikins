package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"

	"epikins-api/internal"
	"epikins-api/internal/controllers/utils"
	"epikins-api/internal/services/projectsService"
)

func ProjectsController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	shouldUpdateProjectList, err := utils.GetQueryBoolValue("update", false, c)
	if err != nil {
		return SendMessage(c, "invalid query parameter", http.StatusBadRequest)
	}

	userLogs, err := utils.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if err != nil {
		return SendMessage(c, "cannot start builds: "+err.Error(), http.StatusInternalServerError)
	}
	projectList, myError := projectsService.ProjectsService(shouldUpdateProjectList, userLogs, appData)
	if myError.Err != nil {
		return SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.JSON(projectList)
}
