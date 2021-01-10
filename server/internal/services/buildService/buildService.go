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
	JobsToBuild []string
	FuMode      bool
	Project     string
	Visibility  libJenkins.Visibility
}

func BuildService(buildParams BuildParams, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	askedProjectData, myError := util.GetLocalProjectData(buildParams.Project, userLogs, appData)
	if myError.Message != "" {
		myError = util.CheckLocalProjectDataError(myError, buildParams.Project, appData.ProjectsCollection)
		return util.GetMyError(BuildError, errors.New(myError.Message), myError.Status)
	}

	err := startBuilds(buildParams, askedProjectData, appData.ProjectsCollection, userLogs)
	if err != nil {
		return util.GetMyError(BuildError, err, http.StatusInternalServerError)
	}
	return internal.MyError{}
}
