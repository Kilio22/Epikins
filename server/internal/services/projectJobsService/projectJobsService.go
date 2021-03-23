package projectJobsService

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

type ProjectJobsParams struct {
	City        string
	Module      string
	ProjectName string
	ForceUpdate bool
}

const ProjectJobsError = "cannot get workgroups associated to given project"

func ProjectJobsService(
	projectJobsParams ProjectJobsParams, userLogs libJenkins.JenkinsCredentials,
	appData *internal.AppData) (
	[]WorkgroupData, internal.MyError,
) {
	localProjectData, myError := util.GetLocalProjectData(projectJobsParams.ProjectName, projectJobsParams.Module, projectJobsParams.ForceUpdate, userLogs, appData)
	if myError.Message != "" {
		return []WorkgroupData{}, util.CheckLocalProjectDataError(myError, projectJobsParams.ProjectName, projectJobsParams.Module, appData.ProjectsCollection)
	}

	workgroups, err := libJenkins.GetWorkgroupsByProject(localProjectData.Job, projectJobsParams.City, userLogs)
	if err != nil {
		return []WorkgroupData{}, util.GetMyError(ProjectJobsError, err, http.StatusInternalServerError)
	} else if len(workgroups) == 0 {
		return []WorkgroupData{}, internal.MyError{}
	}

	workgroupsData, err := getWorkgroupsData(workgroups, localProjectData, projectJobsParams, userLogs, appData.ProjectsCollection)
	if err != nil {
		return []WorkgroupData{}, util.GetMyError(ProjectJobsError, err, http.StatusInternalServerError)
	}
	return workgroupsData, internal.MyError{}
}
