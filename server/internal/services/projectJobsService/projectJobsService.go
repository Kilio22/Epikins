package projectJobsService

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

func ProjectJobsService(projectName string, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) ([]WorkgroupData, internal.MyError) {
	if err := util.CheckProjectsData(userLogs, appData); err != nil {
		return []WorkgroupData{}, internal.MyError{
			Err:        errors.New("cannot get workgroups associated to project \"" + projectName + "\": " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}

	projectsData := appData.ProjectsData[userLogs.Login]
	askedProject, err := util.GetAskedProject(projectsData.ProjectList, projectName)
	if err != nil {
		return []WorkgroupData{}, internal.MyError{
			Err:        errors.New("cannot get workgroups associated to project \"" + projectName + "\": " + err.Error()),
			StatusCode: http.StatusBadRequest,
		}
	}

	workgroups, err := libJenkins.GetWorkgroupsByProject(askedProject.Job, userLogs)
	if err != nil {
		return []WorkgroupData{}, internal.MyError{
			Err:        errors.New("cannot get workgroups associated to project \"" + projectName + "\": " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	} else if len(workgroups) == 0 {
		return []WorkgroupData{}, internal.MyError{Err: nil, StatusCode: http.StatusOK}
	}

	workgroupsData, err := getWorkgroupsData(workgroups, projectName, appData.ProjectsCollection)
	if err != nil {
		return []WorkgroupData{}, internal.MyError{
			Err:        errors.New("cannot get workgroups associated to project \"" + projectName + "\": " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return workgroupsData, internal.MyError{}
}
