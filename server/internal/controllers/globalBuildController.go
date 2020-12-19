package controllers

import (
	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/controllers/utils"
	"epikins-api/internal/services/globalBuildService"
	"epikins-api/internal/services/loginService"
	"epikins-api/pkg/libJenkins"
	"errors"
	"net/http"

	"github.com/gofiber/fiber"
)

func getGlobalBuildParams(c *fiber.Ctx) (globalBuildService.GlobalBuildParams, internal.MyError) {
	visibility, err := utils.GetVisibilityQueryParam(c)
	if err != nil {
		return globalBuildService.GlobalBuildParams{}, internal.MyError{Err: err, StatusCode: http.StatusBadRequest}
	}

	project := c.Query("project")
	if project == "" {
		return globalBuildService.GlobalBuildParams{}, internal.MyError{Err: errors.New("you must specify a project"), StatusCode: http.StatusBadRequest}
	}
	return globalBuildService.GlobalBuildParams{Project: project, Visibility: visibility}, internal.MyError{Err: nil, StatusCode: http.StatusOK}
}

func GlobalBuildController(appData *internal.AppData, c *fiber.Ctx) {
	accessToken := c.Get("Authorization")
	userEmail, err := loginService.LoginService(appData.AppId, accessToken)
	if err != nil {
		sendMessage(c, err.Error(), http.StatusUnauthorized)
		return
	}

	globalBuildParams, myErr := getGlobalBuildParams(c)
	if myErr.Err != nil {
		sendMessage(c, myErr.Err.Error(), myErr.StatusCode)
		return
	}

	userLogs := libJenkins.JenkinsLogs[config.AuthorizedUsers[userEmail]]
	myError := globalBuildService.GlobalBuildService(globalBuildParams, userLogs, appData)
	if myError.Err != nil {
		sendMessage(c, myError.Err.Error(), myError.StatusCode)
		return
	}
	c.SendStatus(http.StatusCreated)
}
