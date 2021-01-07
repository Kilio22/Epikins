package projectsService

import (
	"errors"
	"net/http"
	"time"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

type ProjectResponse struct {
	BuildLimit int            `json:"buildLimit"`
	Job        libJenkins.Job `json:"job"`
	Module     string         `json:"module"`
}

func ProjectsService(shouldUpdateProjectList bool, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) ([]ProjectResponse, internal.MyError) {
	projectsData, ok := appData.ProjectsData[userLogs.Login]
	if ok && !shouldUpdateProjectList && time.Since(projectsData.LastUpdate).Hours() < 1 {
		return getResponseFromProjectList(projectsData.ProjectList, appData.ProjectsCollection)
	}

	if err := util.UpdateLocalProjectList(userLogs, appData); err != nil {
		return []ProjectResponse{}, internal.MyError{
			Err:        errors.New("cannot get projects: " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return getResponseFromProjectList(appData.ProjectsData[userLogs.Login].ProjectList, appData.ProjectsCollection)
}
