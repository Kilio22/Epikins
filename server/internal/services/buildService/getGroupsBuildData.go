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

func getGroupsBuildData(jobs []libJenkins.Job, studentsData []internal.MongoWorkgroupData, project string, collection *mongo.Collection) ([]GroupBuildData, error) {
	var jobsBuildData []GroupBuildData
	for _, job := range jobs {
		if groupMongoData, ok := utils.HasMongoWorkgroupData(job.Name, studentsData); ok {
			jobsBuildData = append(jobsBuildData, GroupBuildData{
				groupJob:       job,
				mongoGroupData: groupMongoData,
			})
		} else {
			newMongoGroupData, err := mongoUtils.AddMongoWorkgroupDataToProject(job, project, collection)
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
