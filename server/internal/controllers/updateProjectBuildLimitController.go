package controllers

import (
	"epikins-api/internal"
	"epikins-api/internal/controllers/utils"
	"epikins-api/internal/services/updateProjectBuildLimitService"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
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
	newLimit, err := getNewLimit(c)
	if err != nil {
		return SendMessage(c, "cannot update project limit: "+err.Error(), http.StatusBadRequest)
	}

	jenkinsCredentials, err := utils.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	myError := updateProjectBuildLimitService.UpdateProjectBuildLimitService(newLimit, projectName, jenkinsCredentials, appData)
	if myError.Err != nil {
		return SendMessage(c, myError.Err.Error(), myError.StatusCode)
	}
	return c.SendStatus(http.StatusCreated)
}
