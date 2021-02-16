package studentJobsService

import (
	"net/http"
	"strings"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

func updateMongoProjectsData(
	mongoProjectsData []internal.MongoProjectData, city string, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) (
	[]internal.MongoProjectData, internal.MyError,
) {
	var definitiveMongoProjectsData []internal.MongoProjectData

	for idx, mongoProjectData := range mongoProjectsData {
		localProjectData, myError := util.GetLocalProjectData(mongoProjectData.Name, mongoProjectData.Module, userLogs, appData)
		if myError.Message != "" {
			if myError.Status == http.StatusBadRequest {
				util.CheckLocalProjectDataError(myError, mongoProjectData.Name, mongoProjectData.Module, appData.ProjectsCollection)
				continue
			}
			return nil, util.CheckLocalProjectDataError(myError, mongoProjectData.Name, mongoProjectData.Module, appData.ProjectsCollection)
		}
		err := util.UpdateMongoProjectData(&mongoProjectsData[idx], localProjectData, city, userLogs, appData.ProjectsCollection)
		if err != nil {
			if strings.Contains(err.Error(), "does not exists on jenkins") {
				continue
			}
			return nil, util.GetMyError(err.Error(), nil, http.StatusInternalServerError)
		}
		definitiveMongoProjectsData = append(definitiveMongoProjectsData, mongoProjectsData[idx])
	}
	return definitiveMongoProjectsData, internal.MyError{}
}
