package controllers

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/globalBuildService"
	"epikins-api/internal/services/util"

	"github.com/gofiber/fiber/v2"
)

func getGlobalBuildParams(c *fiber.Ctx) (globalBuildService.GlobalBuildParams, internal.MyError) {
	visibility, err := controllerUtil.GetVisibilityQueryParam(c)
	if err != nil {
		return globalBuildService.GlobalBuildParams{}, util.GetMyError(globalBuildService.GlobalBuildError, err, http.StatusBadRequest)
	}

	project := c.Query("project")
	if project == "" {
		return globalBuildService.GlobalBuildParams{}, util.GetMyError(globalBuildService.GlobalBuildError+": you must specify a project", nil, http.StatusBadRequest)
	}
	return globalBuildService.GlobalBuildParams{Project: project, Visibility: visibility}, internal.MyError{}
}

func GlobalBuildController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	globalBuildParams, myError := getGlobalBuildParams(c)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}

	userLogs, err := controllerUtil.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(globalBuildService.GlobalBuildError, err, http.StatusInternalServerError), c)
	}
	myError = globalBuildService.GlobalBuildService(globalBuildParams, userLogs, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.SendStatus(http.StatusCreated)
}
