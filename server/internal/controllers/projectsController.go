package controllers

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/controllers/utils"
	"epikins-api/internal/services/loginService"
	"epikins-api/internal/services/projectsService"
	"epikins-api/pkg/libJenkins"
)

func sendProjectList(c *fiber.Ctx, projectList []libJenkins.Job) {
	err := c.JSON(projectList)
	if err != nil {
		log.Println(err)
		c.SendStatus(http.StatusInternalServerError)
	} else {
		c.SendStatus(http.StatusOK)
	}
}

func ProjectsController(appData *internal.AppData, c *fiber.Ctx) {
	accessToken := c.Get("Authorization")
	userEmail, err := loginService.LoginService(appData.AppId, accessToken)
	if err != nil {
		sendMessage(c, err.Error(), http.StatusUnauthorized)
		return
	}

	shouldUpdateProjectList, err := utils.GetQueryBoolValue("update", c)
	if err != nil {
		sendMessage(c, "invalid query parameter", http.StatusBadRequest)
		return
	}

	userLogs := libJenkins.JenkinsLogs[config.AuthorizedUsers[userEmail]]
	projectList, myError := projectsService.ProjectsService(shouldUpdateProjectList, userLogs, appData)
	if myError.Err != nil {
		sendMessage(c, myError.Err.Error(), myError.StatusCode)
		return
	}
	sendProjectList(c, projectList)
}
