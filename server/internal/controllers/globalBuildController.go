package controllers

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/globalBuildService"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

func getGlobalBuildParams(c *fiber.Ctx) (globalBuildService.GlobalBuildParams, internal.MyError) {
	var globalBuildParams globalBuildService.GlobalBuildParams

	err := c.BodyParser(&globalBuildParams)
	if err != nil {
		return globalBuildService.GlobalBuildParams{}, util.GetMyError(globalBuildService.GlobalBuildError, errors.New("wrong body"), http.StatusBadRequest)
	}
	err = validator.New().Struct(globalBuildParams)
	if err != nil {
		return globalBuildService.GlobalBuildParams{}, util.GetMyError(globalBuildService.GlobalBuildError, err, http.StatusBadRequest)
	}
	if globalBuildParams.Visibility != libJenkins.PUBLIC && globalBuildParams.Visibility != libJenkins.PRIVATE {
		return globalBuildService.GlobalBuildParams{}, util.GetMyError(globalBuildService.GlobalBuildError, errors.New("wrong body"), http.StatusBadRequest)
	}
	return globalBuildParams, internal.MyError{}
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
