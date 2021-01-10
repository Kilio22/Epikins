package updateProjectBuildLimitService

import (
	"net/http"
	"time"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

func hasProjectInList(toFind string, projectList []libJenkins.Project) bool {
	for _, project := range projectList {
		if project.Job.Name == toFind {
			return true
		}
	}
	return false
}

func canAddProject(projectName string, jenkinsCredentials libJenkins.JenkinsCredentials, appData *internal.AppData) (
	bool, internal.MyError,
) {
	projectsData, ok := appData.ProjectsData[jenkinsCredentials.Login]
	if ok && time.Since(projectsData.LastUpdate).Hours() < 1 {
		return hasProjectInList(projectName, projectsData.ProjectList), internal.MyError{}
	}
	if err := util.UpdateLocalProjectList(jenkinsCredentials, appData); err != nil {
		return false, util.GetMyError("cannot update project list", err, http.StatusInternalServerError)
	}
	return hasProjectInList(projectName, appData.ProjectsData[jenkinsCredentials.Login].ProjectList), internal.MyError{}
}
