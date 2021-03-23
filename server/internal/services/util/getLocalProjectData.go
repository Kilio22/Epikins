package util

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func GetLocalProjectData(
	projectName string, module string, forceUpdate bool, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) (
	libJenkins.Project, internal.MyError,
) {
	if err := CheckLocalProjectsData(userLogs, forceUpdate, appData); err != nil {
		return libJenkins.Project{}, GetMyError(err.Error(), nil, http.StatusInternalServerError)
	}
	projectsData := appData.ProjectsData[userLogs.Login]
	askedProject, err := GetProjectFromLocalProjectList(projectsData.ProjectList, projectName, module)
	if err != nil {
		return libJenkins.Project{}, GetMyError(err.Error(), nil, http.StatusBadRequest)
	}
	return askedProject, internal.MyError{}
}
