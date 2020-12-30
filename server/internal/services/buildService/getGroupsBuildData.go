package buildService

import (
	"epikins-api/internal"
	"epikins-api/internal/services/utils"
	"epikins-api/internal/services/utils/mongoUtils"
	"epikins-api/pkg/libJenkins"

	"go.mongodb.org/mongo-driver/mongo"
)

type GroupBuildData struct {
	groupJob       libJenkins.Job
	mongoGroupData internal.MongoWorkgroupData
}

// TODO: remove useless workgroups from db
func getGroupsBuildData(jobs []libJenkins.Job, projectData internal.MongoProjectData, project string, collection *mongo.Collection) ([]GroupBuildData, error) {
	var jobsBuildData []GroupBuildData
	for _, job := range jobs {
		if groupMongoData, ok := utils.HasMongoWorkgroupData(job.Name, projectData.MongoWorkgroupsData); ok {
			jobsBuildData = append(jobsBuildData, GroupBuildData{
				groupJob:       job,
				mongoGroupData: groupMongoData,
			})
		} else {
			newMongoGroupData, err := mongoUtils.AddMongoWorkgroupDataToProject(job, project, projectData.BuildLimit, collection)
			if err != nil {
				return []GroupBuildData{}, err
			}
			jobsBuildData = append(jobsBuildData, GroupBuildData{
				groupJob:       job,
				mongoGroupData: newMongoGroupData,
			})
		}
	}
	return jobsBuildData, nil
}
