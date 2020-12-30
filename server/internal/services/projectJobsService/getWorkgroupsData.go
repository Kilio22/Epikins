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

// TODO: remove useless workgroups from db
func getWorkgroupsData(workgroups []libJenkins.Workgroup, project string, collection *mongo.Collection) ([]WorkgroupData, error) {
	projectData, err := utils.FetchProjectData(project, getJobsFromWorkgroups(workgroups), collection)
	if err != nil {
		return []WorkgroupData{}, errors.New("cannot get workgroups remaining builds: " + err.Error())
	}

	var workgroupsData []WorkgroupData
	for _, workgroup := range workgroups {
		if mongoGroupData, ok := utils.HasMongoWorkgroupData(workgroup.Job.Name, projectData.MongoWorkgroupsData); ok {
			workgroupsData = append(workgroupsData, WorkgroupData{
				GroupJob:       workgroup,
				MongoGroupData: mongoGroupData,
			})
		} else {
			newMongoWorkgroupData, err := mongoUtils.AddMongoWorkgroupDataToProject(workgroup.Job, project, projectData.BuildLimit, collection)
			if err != nil {
				return []WorkgroupData{}, err
			}
			workgroupsData = append(workgroupsData, WorkgroupData{
				GroupJob:       workgroup,
				MongoGroupData: newMongoWorkgroupData,
			})
		}
	}
	return workgroupsData, nil
}
