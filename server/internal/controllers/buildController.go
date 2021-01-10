package controllers

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/buildService"
	"epikins-api/internal/services/util"
	"github.com/gofiber/fiber/v2"
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
	visibility, err := controllerUtil.GetVisibilityQueryParam(c)
	if err != nil {
		return buildService.BuildParams{}, util.GetMyError(buildService.BuildError, err, http.StatusBadRequest)
	}

	fuMode, err := controllerUtil.GetQueryBoolValue("fu", false, c)
	if err != nil {
		return buildService.BuildParams{}, util.GetMyError(buildService.BuildError, err, http.StatusBadRequest)
	}

	project := c.Query("project")
	if project == "" {
		return buildService.BuildParams{}, util.GetMyError(buildService.BuildError+": you must specify a project", nil, http.StatusBadRequest)
	}

	jobsToBuild, err := getJobsToBuild(c)
	if err != nil {
		return buildService.BuildParams{}, util.GetMyError(buildService.BuildError+": cannot parse given jobs to build", err, http.StatusBadRequest)
	}
	return buildService.BuildParams{
		JobsToBuild: jobsToBuild,
		FuMode:      fuMode,
		Project:     project,
		Visibility:  visibility,
	}, internal.MyError{}
}

func BuildController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	buildParams, myError := getBuildParams(c)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}

	userLogs, err := controllerUtil.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(buildService.BuildError, err, http.StatusInternalServerError), c)
	}
	myError = buildService.BuildService(buildParams, userLogs, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.SendStatus(http.StatusCreated)
}
