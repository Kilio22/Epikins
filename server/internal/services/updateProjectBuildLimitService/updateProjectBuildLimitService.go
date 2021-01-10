package updateProjectBuildLimitService

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/internal/services/util/mongoUtil"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewLimit struct {
	BuildLimit int `json:"buildLimit" validate:"gte=0"`
}

const UpdateProjectBuildLimitError = "cannot update project build limit"

func checkError(
	err error, shouldAddProject bool, projectName string, userLogs libJenkins.JenkinsCredentials,
	appData *internal.AppData) (bool, internal.MyError) {
	if err == nil {
		return false, internal.MyError{}
	}
	if err == mongo.ErrNoDocuments && shouldAddProject {
		ok, myError := canAddProject(projectName, userLogs, appData)
		if ok {
			return true, internal.MyError{}
		}
		if !ok && myError.Message == "" {
			return false, util.GetMyError(UpdateProjectBuildLimitError+": no project with name \""+projectName+"\" were found", nil, http.StatusBadRequest)
		}
		return false, util.GetMyError(UpdateProjectBuildLimitError, errors.New(myError.Message), myError.Status)
	}
	return false, util.GetMyError(UpdateProjectBuildLimitError, err, http.StatusInternalServerError)
}

func UpdateProjectBuildLimitService(
	newLimit NewLimit, projectName string, userLogs libJenkins.JenkinsCredentials,
	appData *internal.AppData) internal.MyError {
	err := updateProjectData(newLimit, projectName, appData.ProjectsCollection)
	if shouldRetry, myError := checkError(err, true, projectName, userLogs, appData); !shouldRetry || myError.Message != "" {
		return myError
	}

	localProjectData, myError := util.GetLocalProjectData(projectName, userLogs, appData)
	if myError.Message != "" {
		return util.CheckLocalProjectDataError(myError, projectName, appData.ProjectsCollection)
	}

	jobs, err := libJenkins.GetJobsByProject(localProjectData.Job, "REN", userLogs)
	if err != nil {
		return util.GetMyError(UpdateProjectBuildLimitError, err, http.StatusInternalServerError)
	}

	if _, err = mongoUtil.AddMongoProjectData(util.GetNewMongoProjectData(localProjectData, util.GetMongoWorkgroupsDataFromJobs(jobs)), appData.ProjectsCollection); err != nil {
		return util.GetMyError(UpdateProjectBuildLimitError, err, http.StatusInternalServerError)
	}
	err = updateProjectData(newLimit, projectName, appData.ProjectsCollection)
	_, myError = checkError(err, false, projectName, userLogs, appData)
	return myError
}
