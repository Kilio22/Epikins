package projectsService

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

const ProjectsError = "cannot get projects"

type ProjectResponse struct {
	BuildLimit int            `json:"buildLimit"`
	Cities     []string       `json:"cities"`
	Job        libJenkins.Job `json:"job"`
	Module     string         `json:"module"`
}

func ProjectsService(forceUpdate bool, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) (
	[]ProjectResponse, internal.MyError,
) {
	if err := util.CheckLocalProjectsData(userLogs, forceUpdate, appData); err != nil {
		return []ProjectResponse{}, util.GetMyError(ProjectsError, err, http.StatusInternalServerError)
	}

	projectsData := appData.ProjectsData[userLogs.Login]
	return getResponseFromProjectList(projectsData.ProjectList, userLogs, appData.ProjectsCollection)
}
