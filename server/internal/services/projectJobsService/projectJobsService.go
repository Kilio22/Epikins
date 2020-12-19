package projectJobsService

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/utils"
	"epikins-api/pkg/libJenkins"
)

func ProjectJobsService(projectName string, userLogs libJenkins.Logs, appData *internal.AppData) ([]WorkgroupData, internal.MyError) {
	if err := utils.CheckProjectsData(userLogs, appData); err != nil {
		return []WorkgroupData{}, internal.MyError{
			Err:        errors.New("cannot get workgroups associated to project \"" + projectName + "\": " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}

	projectsData := appData.ProjectsData[userLogs.AccountType]
	askedProject, err := utils.GetAskedProject(projectsData.ProjectList, projectName)
	if err != nil {
		return []WorkgroupData{}, internal.MyError{
			Err:        errors.New("cannot get workgroups associated to project \"" + projectName + "\": " + err.Error()),
			StatusCode: http.StatusBadRequest,
		}
	}

	workgroups, err := libJenkins.GetWorkgroupsByProject(askedProject, userLogs)
	if err != nil {
		return []WorkgroupData{}, internal.MyError{
			Err:        errors.New("cannot get workgroups associated to project \"" + projectName + "\": " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	} else if len(workgroups) == 0 {
		return []WorkgroupData{}, internal.MyError{Err: nil, StatusCode: http.StatusOK}
	}

	workgroupsData, err := getWorkgroupsData(workgroups, projectName, appData.Collection)
	if err != nil {
		return []WorkgroupData{}, internal.MyError{
			Err:        errors.New("cannot get workgroups associated to project \"" + projectName + "\": " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return workgroupsData, internal.MyError{Err: nil, StatusCode: http.StatusOK}
}
