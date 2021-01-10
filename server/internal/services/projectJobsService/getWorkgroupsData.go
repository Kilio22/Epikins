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

func getWorkgroupsDataFromMongoProjectData(mongoProjectData internal.MongoProjectData, workgroups []libJenkins.Workgroup) (
	[]WorkgroupData, error,
) {
	var workgroupsData []WorkgroupData
	for _, workgroup := range workgroups {
		if mongoGroupData, ok := util.HasMongoWorkgroupData(workgroup.Job.Name, mongoProjectData.MongoWorkgroupsData); ok {
			workgroupsData = append(workgroupsData, WorkgroupData{
				JobInfos:           workgroup.JobInfos,
				MongoWorkgroupData: mongoGroupData,
			})
		}
	}
	return workgroupsData, nil
}

func getWorkgroupsData(
	workgroups []libJenkins.Workgroup, localProjectData libJenkins.Project, userLogs libJenkins.JenkinsCredentials,
	projectCollection *mongo.Collection) (
	[]WorkgroupData, error,
) {
	mongoProjectData, err := util.GetMongoProjectData(localProjectData, userLogs, projectCollection)
	if err != nil {
		return []WorkgroupData{}, errors.New("cannot get workgroups data: " + err.Error())
	}
	err = util.UpdateMongoProjectData(&mongoProjectData, localProjectData, userLogs, projectCollection)
	if err != nil {
		return []WorkgroupData{}, errors.New("cannot get workgroups data: " + err.Error())
	}
	return getWorkgroupsDataFromMongoProjectData(mongoProjectData, workgroups)
}
