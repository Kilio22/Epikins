package buildService

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

const BuildError = "cannot build"

type BuildParams struct {
	City       string                `json:"city" validate:"required"`
	Jobs       []string              `json:"jobs" validate:"required"`
	Fu         bool                  `json:"fu"`
	Module     string                `json:"module" validate:"required"`
	Project    string                `json:"project" validate:"required"`
	Visibility libJenkins.Visibility `json:"visibility" validate:"required"`
}

func BuildService(buildParams BuildParams, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	askedProjectData, myError := util.GetLocalProjectData(buildParams.Project, buildParams.Module, userLogs, appData)
	if myError.Message != "" {
		myError = util.CheckLocalProjectDataError(myError, buildParams.Project, appData.ProjectsCollection)
		return util.GetMyError(BuildError, errors.New(myError.Message), myError.Status)
	}

	err := startBuilds(buildParams, askedProjectData, buildParams.City, appData.ProjectsCollection, userLogs)
	if err != nil {
		return util.GetMyError(BuildError, err, http.StatusInternalServerError)
	}
	return internal.MyError{}
}
