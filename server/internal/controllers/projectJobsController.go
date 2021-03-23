package controllers

import (
	"errors"
	"net/http"
	"strings"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/projectJobsService"
	"epikins-api/internal/services/util"
	"github.com/gofiber/fiber/v2"
)

func getProjectJobsParams(c *fiber.Ctx) (projectJobsService.ProjectJobsParams, error) {
	projectName := c.Params("project")
	module := c.Params("module")
	city := strings.ToUpper(c.Params("city"))
	shouldUpdateProjectList, err := controllerUtil.GetQueryBoolValue("update", false, c)
	if err != nil {
		return projectJobsService.ProjectJobsParams{}, err
	}
	return projectJobsService.ProjectJobsParams{
		City:        city,
		Module:      module,
		ProjectName: projectName,
		ForceUpdate: shouldUpdateProjectList,
	}, nil
}

func ProjectJobsController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	projectJobsParams, err := getProjectJobsParams(c)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(projectJobsService.ProjectJobsError, errors.New("invalid query parameter"), http.StatusBadRequest), c)
	}

	userLogs, err := controllerUtil.GetUserJenkinsCredentials(userEmail, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(projectJobsService.ProjectJobsError, err, http.StatusInternalServerError), c)
	}

	workgroupsData, myError := projectJobsService.ProjectJobsService(projectJobsParams, userLogs, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.JSON(workgroupsData)
}
