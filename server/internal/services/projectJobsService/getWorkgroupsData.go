package projectJobsService

import (
	"errors"

	"epikins-api/internal"
	"epikins-api/internal/services/utils"
	"epikins-api/internal/services/utils/mongoUtils"
	"epikins-api/pkg/libJenkins"

	"go.mongodb.org/mongo-driver/mongo"
)

type WorkgroupData struct {
	GroupJob       libJenkins.Workgroup        `json:"groupJob"`
	MongoGroupData internal.MongoWorkgroupData `json:"mongoGroupData"`
}

func getWorkgroupsData(workgroups []libJenkins.Workgroup, project string, collection *mongo.Collection) ([]WorkgroupData, error) {
	projectData, err := mongoUtils.FetchProjectData(project, getJobsFromWorkgroups(workgroups), collection)
	if err != nil {
		return []WorkgroupData{}, errors.New("cannot get workgroups remaining builds: " + err.Error())
	}

	var workgroupsData []WorkgroupData
	for _, workgroup := range workgroups {
		if mongoGroupData, ok := utils.HasMongoWorkgoupData(workgroup.Job.Name, projectData.MongoWorkgroupsData); ok {
			workgroupsData = append(workgroupsData, WorkgroupData{
				GroupJob:       workgroup,
				MongoGroupData: mongoGroupData,
			})
		} else {
			newMongoworkgroupData, err := mongoUtils.AddMongoWorkgroupDataToProject(workgroup.Job, project, collection)
			if err != nil {
				return []WorkgroupData{}, err
			}
			workgroupsData = append(workgroupsData, WorkgroupData{
				GroupJob:       workgroup,
				MongoGroupData: newMongoworkgroupData,
			})
		}
	}
	return workgroupsData, nil
}
