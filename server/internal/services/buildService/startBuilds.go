package buildService

import (
	"errors"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/internal/services/util/mongoUtil"
	"go.mongodb.org/mongo-driver/bson"

	"epikins-api/pkg/libJenkins"

	"go.mongodb.org/mongo-driver/mongo"
)

const StartBuildsError = "cannot start builds"

func buildLoop(buildParams BuildParams, groupsBuildData []internal.MongoWorkgroupData, userLogs libJenkins.JenkinsCredentials) error {
	for _, jobName := range buildParams.JobsToBuild {
		for idx := range groupsBuildData {
			if groupsBuildData[idx].Name != jobName {
				continue
			}
			shouldBreak, err := startBuild(&groupsBuildData[idx], buildParams, userLogs)
			if err != nil {
				return err
			}
			if shouldBreak {
				break
			}
		}
	}
	return nil
}

func startBuilds(
	buildParams BuildParams, localProjectData libJenkins.Project, projectCollection *mongo.Collection,
	userLogs libJenkins.JenkinsCredentials,
) error {
	mongoProjectData, err := util.GetMongoProjectData(localProjectData, userLogs, projectCollection)
	if err != nil {
		return errors.New(StartBuildsError + err.Error())
	}
	err = util.UpdateMongoProjectData(&mongoProjectData, localProjectData, userLogs, projectCollection)
	if err != nil {
		return errors.New(StartBuildsError + ": cannot get workgroups data: " + err.Error())
	}

	err = buildLoop(buildParams, mongoProjectData.MongoWorkgroupsData, userLogs)
	if err != nil {
		return err
	}
	return mongoUtil.UpdateProject(mongoProjectData.Name,
		bson.M{"$set": bson.M{"mongoworkgroupsdata": mongoProjectData.MongoWorkgroupsData}}, projectCollection)
}
