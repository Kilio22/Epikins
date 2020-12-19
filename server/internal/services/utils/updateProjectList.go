package utils

import (
	"time"

	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func UpdateProjectList(userLogs libJenkins.Logs, appData *internal.AppData) error {
	projectList, err := libJenkins.GetProjects(userLogs)

	if err != nil {
		return err
	}
	appData.ProjectsData[userLogs.AccountType] = internal.ProjectsData{ProjectList: projectList, LastUpdate: time.Now()}
	return nil
}
