package projectsService

import (
	"errors"
	"net/http"
	"time"

	"epikins-api/internal"
	"epikins-api/internal/services/utils"
	"epikins-api/pkg/libJenkins"
)

func ProjectsService(shouldUpdateProjectList bool, userLogs libJenkins.Logs, appData *internal.AppData) ([]libJenkins.Job, internal.MyError) {
	projectsData, ok := appData.ProjectsData[userLogs.AccountType]
	if ok && !shouldUpdateProjectList && time.Since(projectsData.LastUpdate).Hours() < 1 {
		return projectsData.ProjectList, internal.MyError{
			Err:        nil,
			StatusCode: http.StatusOK,
		}
	}

	if err := utils.UpdateProjectList(userLogs, appData); err != nil {
		return []libJenkins.Job{}, internal.MyError{
			Err:        errors.New("cannot get projects: " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return appData.ProjectsData[userLogs.AccountType].ProjectList, internal.MyError{}
}
