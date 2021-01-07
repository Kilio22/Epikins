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

func getWorkgroupsData(workgroups []libJenkins.Workgroup, project libJenkins.Project, collection *mongo.Collection) (
	[]WorkgroupData, error,
) {
	jobs := getJobsFromWorkgroups(workgroups)
	mongoProjectData, err := util.GetMongoProjectData(project, jobs, collection)
	if err != nil {
		return []WorkgroupData{}, errors.New("cannot get workgroups data: " + err.Error())
	}
	err = util.UpdateMongoProjectData(&mongoProjectData, jobs, collection)
	if err != nil {
		return []WorkgroupData{}, errors.New("cannot get workgroups data: " + err.Error())
	}
	return getWorkgroupsDataFromMongoProjectData(mongoProjectData, workgroups)
}
