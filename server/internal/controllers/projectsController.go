package controllers

import (
	"epikins-api/internal/services/loginService"
	"net/http"

	"github.com/gofiber/fiber/v2"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/controllers/utils"
	"epikins-api/internal/services/projectsService"
	"epikins-api/pkg/libJenkins"
)

func ProjectsController(appData *internal.AppData, c *fiber.Ctx) error {
	accessToken := c.Get("Authorization")
	userEmail, err := loginService.LoginService(appData.AppId, accessToken)
	if err != nil {
		return SendMessage(c, err.Error(), http.StatusUnauthorized)
	}

	shouldUpdateProjectList, err := utils.GetQueryBoolValue("update", false, c)
	if err != nil {
		return SendMessage(c, "invalid query parameter", http.StatusBadRequest)
	}

	userLogs := libJenkins.JenkinsLogs[config.AuthorizedUsers[userEmail]]
	projectList, myError := projectsService.ProjectsService(shouldUpdateProjectList, userLogs, appData)
	if myError.Err != nil {
		return SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.JSON(projectList)
}
