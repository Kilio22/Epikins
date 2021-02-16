package controllers

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/updateProjectBuildLimitService"
	"epikins-api/internal/services/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func getNewLimit(c *fiber.Ctx) (updateProjectBuildLimitService.NewLimit, error) {
	var newLimit updateProjectBuildLimitService.NewLimit
	err := c.BodyParser(&newLimit)
	if err != nil {
		return updateProjectBuildLimitService.NewLimit{}, err
	}
	err = validator.New().Struct(newLimit)
	if err != nil {
		return updateProjectBuildLimitService.NewLimit{}, err
	}
	return newLimit, nil
}

func UpdateProjectBuildLimitController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	projectName := c.Params("project")
	module := c.Params("module")
	newLimit, err := getNewLimit(c)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(updateProjectBuildLimitService.UpdateProjectBuildLimitError, err, http.StatusBadRequest), c)
	}

	jenkinsCredentials, err := controllerUtil.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(updateProjectBuildLimitService.UpdateProjectBuildLimitError, err, http.StatusInternalServerError), c)
	}

	myError := updateProjectBuildLimitService.UpdateProjectBuildLimitService(newLimit, projectName, module, jenkinsCredentials, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.SendStatus(http.StatusCreated)
}
