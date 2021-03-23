package globalBuildService

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

const GlobalBuildError = "cannot build"

type GlobalBuildParams struct {
	City       string                `json:"city" validate:"required"`
	Module     string                `json:"module" validate:"required"`
	Project    string                `json:"project" validate:"required"`
	Visibility libJenkins.Visibility `json:"visibility" validate:"required"`
}

func GlobalBuildService(
	globalBuildParams GlobalBuildParams, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	askedProjectData, myError := util.GetLocalProjectData(globalBuildParams.Project, globalBuildParams.Module, false, userLogs, appData)
	if myError.Message != "" {
		return util.CheckLocalProjectDataError(myError, globalBuildParams.Project, globalBuildParams.Module, appData.ProjectsCollection)
	}

	globalJobUrl, err := libJenkins.GetGlobalJobUrlByProject(askedProjectData.Job, globalBuildParams.City, userLogs)
	if err != nil {
		return util.GetMyError(GlobalBuildError, err, http.StatusInternalServerError)
	}

	err = libJenkins.BuildJob(globalJobUrl, globalBuildParams.Visibility, userLogs)
	if err != nil {
		return util.GetMyError(GlobalBuildError, err, http.StatusInternalServerError)
	}
	return internal.MyError{}
}
