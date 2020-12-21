package buildService

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/utils"
	"epikins-api/pkg/libJenkins"
)

type BuildParams struct {
	JobsToBuild []string
	FuMode      bool
	Project     string
	Visibility  libJenkins.Visibility
}

func BuildService(buildParams BuildParams, appData *internal.AppData, userLogs libJenkins.Logs) internal.MyError {
	if err := utils.CheckProjectsData(userLogs, appData); err != nil {
		return internal.MyError{Err: errors.New("cannot build: " + err.Error()), StatusCode: http.StatusInternalServerError}
	}

	projectsData := appData.ProjectsData[userLogs.AccountType]
	askedProjectData, err := utils.GetAskedProject(projectsData.ProjectList, buildParams.Project)
	if err != nil {
		return internal.MyError{
			Err:        errors.New("cannot build: " + err.Error()),
			StatusCode: http.StatusBadRequest,
		}
	}

	jobs, err := libJenkins.GetJobsByProject(askedProjectData, userLogs)
	if err != nil {
		return internal.MyError{Err: errors.New("cannot build: " + err.Error()), StatusCode: http.StatusInternalServerError}
	} else if len(jobs) == 0 {
		return internal.MyError{Err: errors.New("cannot build: no jobs to build for this project"), StatusCode: http.StatusBadRequest}
	}

	err = startBuilds(buildParams, jobs, appData.ProjectsCollection, userLogs)
	if err != nil {
		return internal.MyError{Err: errors.New("cannot build: " + err.Error()), StatusCode: http.StatusInternalServerError}
	}
	return internal.MyError{}
}
