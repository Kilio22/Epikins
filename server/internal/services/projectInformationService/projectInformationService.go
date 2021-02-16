package projectInformationService

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/projectsService"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

const ProjectInformationError = "cannot retrieve cities linked to the given project"

func ProjectInformationService(projectName string, module string, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) (
	projectsService.ProjectResponse, internal.MyError,
) {
	localProjectData, myError := util.GetLocalProjectData(projectName, module, userLogs, appData)
	if myError.Message != "" {
		return projectsService.ProjectResponse{}, util.CheckLocalProjectDataError(myError, projectName, module, appData.ProjectsCollection)
	}

	mongoProjectData, err := util.GetMongoProjectData(localProjectData, "", userLogs, appData.ProjectsCollection)
	if err != nil {
		return projectsService.ProjectResponse{}, util.GetMyError(ProjectInformationError, err, http.StatusInternalServerError)
	}
	return projectsService.ProjectResponse{
		BuildLimit: mongoProjectData.BuildLimit,
		Cities:     localProjectData.Cities,
		Job:        localProjectData.Job,
		Module:     localProjectData.Module,
	}, internal.MyError{}
}
