package globalBuildService

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

const GlobalBuildError = "cannot build"

type GlobalBuildParams struct {
	Project    string
	Visibility libJenkins.Visibility
}

func GlobalBuildService(
	globalBuildParams GlobalBuildParams, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	askedProjectData, myError := util.GetLocalProjectData(globalBuildParams.Project, userLogs, appData)
	if myError.Message != "" {
		return util.CheckLocalProjectDataError(myError, globalBuildParams.Project, appData.ProjectsCollection)
	}

	globalJobUrl, err := libJenkins.GetGlobalJobUrlByProject(askedProjectData.Job, "REN", userLogs)
	if err != nil {
		return util.GetMyError(GlobalBuildError, err, http.StatusInternalServerError)
	}

	err = libJenkins.BuildJob(globalJobUrl, globalBuildParams.Visibility, userLogs)
	if err != nil {
		return util.GetMyError(GlobalBuildError, err, http.StatusInternalServerError)
	}
	return internal.MyError{}
}
