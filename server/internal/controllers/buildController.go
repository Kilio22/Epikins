package controllers

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/buildService"
	"epikins-api/internal/services/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func getBuildParams(c *fiber.Ctx) (buildService.BuildParams, internal.MyError) {
	var buildParams buildService.BuildParams

	err := c.BodyParser(&buildParams)
	if err != nil {
		return buildService.BuildParams{}, util.GetMyError(buildService.BuildError, errors.New("wrong body"), http.StatusBadRequest)
	}
	err = validator.New().Struct(buildParams)
	if err != nil {
		return buildService.BuildParams{}, util.GetMyError(buildService.BuildError, err, http.StatusBadRequest)
	}
	return buildParams, internal.MyError{}
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
