package utils

import (
	"time"

	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func CheckProjectsData(userLogs libJenkins.Logs, appData *internal.AppData) error {
	projectsData, ok := appData.ProjectsData[userLogs.AccountType]
	if !ok || time.Since(projectsData.LastUpdate).Hours() > 1 {
		if err := UpdateProjectList(userLogs, appData); err != nil {
			return err
		}
	}
	return nil
}
