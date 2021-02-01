package util

import (
	"time"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func CheckLocalProjectsData(userLogs libJenkins.JenkinsCredentials, forceUpdate bool, appData *internal.AppData) error {
	projectsData, ok := appData.ProjectsData[userLogs.Login]
	if !ok || time.Since(projectsData.LastUpdate).Hours() > config.LocalProjectListRefreshTime || forceUpdate {
		if err := UpdateLocalProjectList(userLogs, appData); err != nil {
			return err
		}
	}
	return nil
}
