package controllers

import (
	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/services/loginService"
	"epikins-api/internal/services/projectJobsService"
	"epikins-api/pkg/libJenkins"
	"log"
	"net/http"

	"github.com/gofiber/fiber"
)

func ProjectJobsController(appData *internal.AppData, c *fiber.Ctx) {
	accessToken := c.Get("Authorization")
	userEmail, err := loginService.LoginService(appData.AppId, accessToken)
	if err != nil {
		sendMessage(c, err.Error(), http.StatusUnauthorized)
		return
	}

	projectName := c.Params("project")
	userLogs := libJenkins.JenkinsLogs[config.AuthorizedUsers[userEmail]]
	groupsData, myError := projectJobsService.ProjectJobsService(projectName, userLogs, appData)
	if myError.Err != nil {
		sendMessage(c, myError.Err.Error(), myError.StatusCode)
		return
	}
	err = c.JSON(groupsData)
	if err != nil {
		log.Println(err)
		c.SendStatus(http.StatusInternalServerError)
		return
	}
	c.SendStatus(http.StatusOK)
}
