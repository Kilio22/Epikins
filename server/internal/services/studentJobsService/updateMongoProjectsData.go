package studentJobsService

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

func updateMongoProjectsData(
	mongoProjectsData []internal.MongoProjectData, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) (
	[]internal.MongoProjectData, internal.MyError,
) {
	var definitiveMongoProjectsData []internal.MongoProjectData

	for idx, mongoProjectData := range mongoProjectsData {
		localProjectData, myError := util.GetLocalProjectData(mongoProjectData.Name, userLogs, appData)
		if myError.Message != "" {
			if myError.Status == http.StatusBadRequest {
				util.CheckLocalProjectDataError(myError, mongoProjectData.Name, appData.ProjectsCollection)
				continue
			}
			return nil, util.CheckLocalProjectDataError(myError, mongoProjectData.Name, appData.ProjectsCollection)
		}
		err := util.UpdateMongoProjectData(&mongoProjectsData[idx], localProjectData, userLogs, appData.ProjectsCollection)
		if err != nil {
			return nil, util.GetMyError(err.Error(), nil, http.StatusInternalServerError)
		}
		definitiveMongoProjectsData = append(definitiveMongoProjectsData, mongoProjectsData[idx])
	}
	return definitiveMongoProjectsData, internal.MyError{}
}
