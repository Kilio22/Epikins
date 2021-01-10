package controllers

import (
	"net/http"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/studentJobsService"
	"epikins-api/internal/services/util"
	"github.com/gofiber/fiber/v2"
)

func StudentJobsController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	userLogs, err := controllerUtil.GetJenkinsCredentials(config.HighestPrivilegeJenkinsLogin, appData.JenkinsCredentialsCollection)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(studentJobsService.StudentJobsError, err, http.StatusInternalServerError), c)
	}

	studentJobs, myError := studentJobsService.StudentJobsService(userEmail, userLogs, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.JSON(studentJobs)
}
