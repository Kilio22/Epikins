package controllers

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/buildService"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func getBuildInfo(c *fiber.Ctx, userEmail string) (buildService.BuildInfo, internal.MyError) {
	var buildParams buildService.BuildParams

	err := c.BodyParser(&buildParams)
	if err != nil {
		return buildService.BuildInfo{}, util.GetMyError(buildService.BuildError, errors.New("wrong body"), http.StatusBadRequest)
	}
	err = validator.New().Struct(buildParams)
	if err != nil {
		return buildService.BuildInfo{}, util.GetMyError(buildService.BuildError, err, http.StatusBadRequest)
	}
	if buildParams.Visibility != libJenkins.PUBLIC && buildParams.Visibility != libJenkins.PRIVATE {
		return buildService.BuildInfo{}, util.GetMyError(buildService.BuildError, errors.New("wrong body"), http.StatusBadRequest)
	}
	return buildService.BuildInfo{
		BuildParams: buildParams,
		Starter:     userEmail,
	}, internal.MyError{}
}

func BuildController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	buildInfo, myError := getBuildInfo(c, userEmail)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}

	userLogs, err := controllerUtil.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(buildService.BuildError, err, http.StatusInternalServerError), c)
	}
	myError = buildService.BuildService(buildInfo, userLogs, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.SendStatus(http.StatusCreated)
}
