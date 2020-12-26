package utils

import (
	"time"

	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func CheckProjectsData(userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) error {
	projectsData, ok := appData.ProjectsData[userLogs.Login]
	if !ok || time.Since(projectsData.LastUpdate).Hours() > 1 {
		if err := UpdateProjectList(userLogs, appData); err != nil {
			return err
		}
	}
	return nil
}
