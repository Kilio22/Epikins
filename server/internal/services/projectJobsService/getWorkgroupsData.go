package projectJobsService

import (
	"errors"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"

	"go.mongodb.org/mongo-driver/mongo"
)

type WorkgroupData struct {
	JobInfos           libJenkins.JobInfos         `json:"jobInfos"`
	MongoWorkgroupData internal.MongoWorkgroupData `json:"mongoWorkgroupData"`
}

func getWorkgroupsDataFromMongoProjectData(mongoProjectData internal.MongoProjectData, city string, workgroups []libJenkins.Workgroup) (
	[]WorkgroupData, error,
) {
	var workgroupsData []WorkgroupData
	for _, workgroup := range workgroups {
		if mongoGroupData, ok := util.HasMongoWorkgroupData(workgroup.Job.Name, mongoProjectData.CitiesData[city].MongoWorkgroupsData); ok {
			workgroupsData = append(workgroupsData, WorkgroupData{
				JobInfos:           workgroup.JobInfos,
				MongoWorkgroupData: mongoGroupData,
			})
		}
	}
	return workgroupsData, nil
}

func getWorkgroupsData(
	workgroups []libJenkins.Workgroup, localProjectData libJenkins.Project, projectJobsParams ProjectJobsParams,
	userLogs libJenkins.JenkinsCredentials,
	projectCollection *mongo.Collection) (
	[]WorkgroupData, error,
) {
	mongoProjectData, err := util.GetMongoProjectData(localProjectData, projectJobsParams.City, userLogs, projectCollection)
	if err != nil {
		return []WorkgroupData{}, errors.New("cannot get workgroups data: " + err.Error())
	}
	err = util.UpdateMongoProjectData(&mongoProjectData, localProjectData, projectJobsParams.City, projectJobsParams.ForceUpdate, userLogs, projectCollection)
	if err != nil {
		return []WorkgroupData{}, errors.New("cannot get workgroups data: " + err.Error())
	}
	return getWorkgroupsDataFromMongoProjectData(mongoProjectData, projectJobsParams.City, workgroups)
}
