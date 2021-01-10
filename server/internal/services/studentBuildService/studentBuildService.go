package studentBuildService

import (
	"errors"
	"net/http"
	"strings"

	"epikins-api/internal"
	"epikins-api/internal/services/buildService"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

type StudentBuildParams struct {
	Group   string `json:"group" validate:"required"`
	Project string `json:"project" validate:"required"`
}

const StudentBuildError = "cannot build"

func StudentBuildService(
	studentEmail string, studentBuildParams StudentBuildParams, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	studentName := util.GetUsernameFromEmail(studentEmail)
	if !strings.Contains(studentBuildParams.Group, studentName) {
		return util.GetMyError(StudentBuildError, errors.New("you can't start a build for another group"), http.StatusBadRequest)
	}

	buildParams := buildService.BuildParams{
		JobsToBuild: []string{studentBuildParams.Group},
		FuMode:      false,
		Project:     studentBuildParams.Project,
		Visibility:  libJenkins.PUBLIC,
	}
	return buildService.BuildService(buildParams, userLogs, appData)
}
