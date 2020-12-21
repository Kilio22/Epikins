package controllers

import (
	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/services/projectJobsService"
	"epikins-api/pkg/libJenkins"
	"github.com/gofiber/fiber/v2"
)

func ProjectJobsController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	projectName := c.Params("project")
	userLogs := libJenkins.JenkinsLogs[config.AuthorizedUsers[userEmail]]
	groupsData, myError := projectJobsService.ProjectJobsService(projectName, userLogs, appData)
	if myError.Err != nil {
		return SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.JSON(groupsData)
}
