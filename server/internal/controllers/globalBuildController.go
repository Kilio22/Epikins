package controllers

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers/util"
	"epikins-api/internal/services/globalBuildService"
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func getGlobalBuildParams(c *fiber.Ctx) (globalBuildService.GlobalBuildParams, internal.MyError) {
	visibility, err := util.GetVisibilityQueryParam(c)
	if err != nil {
		return globalBuildService.GlobalBuildParams{}, internal.MyError{Err: err, StatusCode: http.StatusBadRequest}
	}

	project := c.Query("project")
	if project == "" {
		return globalBuildService.GlobalBuildParams{}, internal.MyError{Err: errors.New("you must specify a project"), StatusCode: http.StatusBadRequest}
	}
	return globalBuildService.GlobalBuildParams{Project: project, Visibility: visibility}, internal.MyError{Err: nil, StatusCode: http.StatusOK}
}

func GlobalBuildController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	globalBuildParams, myErr := getGlobalBuildParams(c)
	if myErr.Err != nil {
		return SendMessage(c, myErr.Err.Error(), myErr.StatusCode)
	}

	userLogs, err := util.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if err != nil {
		return SendMessage(c, "cannot start build: "+err.Error(), http.StatusInternalServerError)
	}
	myError := globalBuildService.GlobalBuildService(globalBuildParams, userLogs, appData)
	if myError.Err != nil {
		return SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.SendStatus(http.StatusCreated)
}
