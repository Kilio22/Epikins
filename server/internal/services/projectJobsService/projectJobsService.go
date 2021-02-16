package projectJobsService

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

const ProjectJobsError = "cannot get workgroups associated to given project"

func ProjectJobsService(projectName string, module string, city string, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) (
	[]WorkgroupData, internal.MyError,
) {
	localProjectData, myError := util.GetLocalProjectData(projectName, module, userLogs, appData)
	if myError.Message != "" {
		return []WorkgroupData{}, util.CheckLocalProjectDataError(myError, projectName, module, appData.ProjectsCollection)
	}

	workgroups, err := libJenkins.GetWorkgroupsByProject(localProjectData.Job, city, userLogs)
	if err != nil {
		return []WorkgroupData{}, util.GetMyError(ProjectJobsError, err, http.StatusInternalServerError)
	} else if len(workgroups) == 0 {
		return []WorkgroupData{}, internal.MyError{}
	}

	workgroupsData, err := getWorkgroupsData(workgroups, localProjectData, city, userLogs, appData.ProjectsCollection)
	if err != nil {
		return []WorkgroupData{}, util.GetMyError(ProjectJobsError, err, http.StatusInternalServerError)
	}
	return workgroupsData, internal.MyError{}
}
