package controllers

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/controllers/controllerUtil"
	"epikins-api/internal/services/studentBuildService"
	"epikins-api/internal/services/util"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var StudentJenkinsLogin = util.GetEnvVariable("STUDENT_JENKINS_LOGIN")

func getStudentBuildParams(c *fiber.Ctx) (studentBuildService.StudentBuildParams, error) {
	var studentBuildParams studentBuildService.StudentBuildParams

	err := c.BodyParser(&studentBuildParams)
	if err != nil {
		return studentBuildService.StudentBuildParams{}, errors.New("wrong body")
	}
	err = validator.New().Struct(studentBuildParams)
	if err != nil {
		return studentBuildService.StudentBuildParams{}, errors.New("wrong body")
	}
	return studentBuildParams, nil
}

func StudentBuildController(appData *internal.AppData, c *fiber.Ctx) error {
	userEmail := c.Get("email")
	studentBuildParams, err := getStudentBuildParams(c)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(studentBuildService.StudentBuildError, err, http.StatusInternalServerError), c)
	}

	userLogs, err := controllerUtil.GetJenkinsCredentials(StudentJenkinsLogin, appData.JenkinsCredentialsCollection)
	if err != nil {
		return controllerUtil.SendMyError(util.GetMyError(studentBuildService.StudentBuildError, err, http.StatusInternalServerError), c)
	}

	myError := studentBuildService.StudentBuildService(userEmail, studentBuildParams, userLogs, appData)
	if myError.Message != "" {
		return controllerUtil.SendMyError(myError, c)
	}
	return c.SendStatus(http.StatusCreated)
}
