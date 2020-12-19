package controllers

import (
	"errors"
	"net/http"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/controllers/utils"
	"epikins-api/internal/services/buildService"
	"epikins-api/internal/services/loginService"
	"epikins-api/pkg/libJenkins"

	"github.com/gofiber/fiber"
)

func getJobsToBuild(c *fiber.Ctx) ([]string, error) {
	var jobsToBuild []string
	err := c.BodyParser(&jobsToBuild)
	if err != nil {
		return []string{}, err
	}
	return jobsToBuild, nil
}

func getBuildParams(c *fiber.Ctx) (buildService.BuildParams, internal.MyError) {
	visibility, err := utils.GetVisibilityQueryParam(c)
	if err != nil {
		return buildService.BuildParams{}, internal.MyError{Err: err, StatusCode: http.StatusBadRequest}
	}

	fuMode, err := utils.GetQueryBoolValue("fu", c)
	if err != nil {
		return buildService.BuildParams{}, internal.MyError{Err: err, StatusCode: http.StatusBadRequest}
	}

	project := c.Query("project")
	if project == "" {
		return buildService.BuildParams{}, internal.MyError{Err: errors.New("you must specify a project"), StatusCode: http.StatusBadRequest}
	}

	jobsToBuild, err := getJobsToBuild(c)
	if err != nil {
		return buildService.BuildParams{}, internal.MyError{
			Err:        errors.New("cannot parse given jobs to build: " + err.Error()),
			StatusCode: http.StatusBadRequest,
		}
	}
	return buildService.BuildParams{
		JobsToBuild: jobsToBuild,
		FuMode:      fuMode,
		Project:     project,
		Visibility:  visibility,
	}, internal.MyError{Err: nil, StatusCode: http.StatusOK}
}

func BuildController(appData *internal.AppData, c *fiber.Ctx) {
	accessToken := c.Get("Authorization")
	userEmail, err := loginService.LoginService(appData.AppId, accessToken)
	if err != nil {
		sendMessage(c, err.Error(), http.StatusUnauthorized)
		return
	}

	buildParams, myErr := getBuildParams(c)
	if myErr.Err != nil {
		sendMessage(c, myErr.Err.Error(), myErr.StatusCode)
		return
	}

	userLogs := libJenkins.JenkinsLogs[config.AuthorizedUsers[userEmail]]
	myError := buildService.BuildService(buildParams, appData, userLogs)
	if myError.Err != nil {
		sendMessage(c, myError.Err.Error(), myError.StatusCode)
		return
	}
	c.SendStatus(http.StatusCreated)
}
