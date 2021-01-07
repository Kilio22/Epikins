package util

import (
	"time"

	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func UpdateLocalProjectList(userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) error {
	projectList, err := libJenkins.GetProjects(userLogs)

	if err != nil {
		return err
	}
	appData.ProjectsData[userLogs.Login] = internal.ProjectsData{ProjectList: projectList, LastUpdate: time.Now()}
	return nil
}
