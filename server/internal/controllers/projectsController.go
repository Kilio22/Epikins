package controllers

import (
	"errors"
	"net/http"

	"epikins-api/internal/services/util"
	"github.com/gofiber/fiber/v2"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/projectsService"
)

func ProjectsController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	city := c.Params("city")
	shouldUpdateProjectList, err := controllerUtil.GetQueryBoolValue("update", false, c)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(projectsService.ProjectsError, errors.New("invalid query parameter"), http.StatusBadRequest), c)
	}

	userLogs, err := controllerUtil.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(projectsService.ProjectsError, err, http.StatusInternalServerError), c)
	}
	projectList, myError := projectsService.ProjectsService(shouldUpdateProjectList, city, userLogs, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.JSON(projectList)
}
