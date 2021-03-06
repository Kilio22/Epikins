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
	City    string `json:"city"  validate:"required"`
	Group   string `json:"group" validate:"required"`
	Module  string `json:"module" validate:"required"`
	Project string `json:"project" validate:"required"`
}

const StudentBuildError = "cannot build"

func StudentBuildService(
	studentEmail string, studentBuildParams StudentBuildParams, userLogs libJenkins.JenkinsCredentials,
	appData *internal.AppData) internal.MyError {
	studentName := util.GetUsernameFromEmail(studentEmail)
	if !strings.Contains(studentBuildParams.Group, studentName) {
		return util.GetMyError(StudentBuildError, errors.New("you can't start a build for another group"), http.StatusBadRequest)
	}

	buildParams := buildService.BuildParams{
		City: studentBuildParams.City,
		Jobs: []string{
			studentBuildParams.Group,
		},
		Fu:         false,
		Module:     studentBuildParams.Module,
		Project:    studentBuildParams.Project,
		Visibility: libJenkins.PUBLIC,
	}
	buildInfo := buildService.BuildInfo{BuildParams: buildParams, Starter: studentEmail}
	return buildService.BuildService(buildInfo, userLogs, appData)
}
