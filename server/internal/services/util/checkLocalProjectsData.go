package util

import (
	"time"

	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func CheckLocalProjectsData(userLogs libJenkins.JenkinsCredentials, forceUpdate bool, appData *internal.AppData) error {
	projectsData, ok := appData.ProjectsData[userLogs.Login]
	if !ok || time.Since(projectsData.LastUpdate).Hours() > 4 || forceUpdate {
		if err := UpdateLocalProjectList(userLogs, appData); err != nil {
			return err
		}
	}
	return nil
}
