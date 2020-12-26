package globalBuildService

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/utils"
	"epikins-api/pkg/libJenkins"
)

type GlobalBuildParams struct {
	Project    string
	Visibility libJenkins.Visibility
}

func GlobalBuildService(globalBuildParams GlobalBuildParams, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	if err := utils.CheckProjectsData(userLogs, appData); err != nil {
		return internal.MyError{Err: errors.New("cannot build: " + err.Error()), StatusCode: http.StatusInternalServerError}
	}

	projectsData := appData.ProjectsData[userLogs.Login]
	askedProject, err := utils.GetAskedProject(projectsData.ProjectList, globalBuildParams.Project)
	if err != nil {
		return internal.MyError{Err: errors.New("cannot build: " + err.Error()), StatusCode: http.StatusBadRequest}
	}

	globalJobUrl, err := libJenkins.GetGlobalJobUrlByProject(askedProject, userLogs)
	if err != nil {
		return internal.MyError{Err: errors.New("cannot build: " + err.Error()), StatusCode: http.StatusInternalServerError}
	}

	err = libJenkins.BuildJob(globalJobUrl, globalBuildParams.Visibility, userLogs)
	if err != nil {
		return internal.MyError{Err: errors.New("cannot build: " + err.Error()), StatusCode: http.StatusInternalServerError}
	}
	return internal.MyError{}
}
