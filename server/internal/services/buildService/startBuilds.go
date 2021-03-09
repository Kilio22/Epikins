package buildService

import (
	"errors"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/internal/services/util/mongoUtil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"epikins-api/pkg/libJenkins"
)

const StartBuildsError = "cannot start builds"

func buildLoop(
	buildInfo BuildInfo, groupsBuildData []internal.MongoWorkgroupData, userLogs libJenkins.JenkinsCredentials,
	buildLogsCollection *mongo.Collection) error {
	for _, jobName := range buildInfo.BuildParams.Jobs {
		for idx := range groupsBuildData {
			if groupsBuildData[idx].Name != jobName {
				continue
			}
			shouldBreak, err := startBuild(&groupsBuildData[idx], buildInfo, userLogs, buildLogsCollection)
			if err != nil && !shouldBreak {
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
	buildInfo BuildInfo, localProjectData libJenkins.Project, appData *internal.AppData,
	userLogs libJenkins.JenkinsCredentials,
) error {
	mongoProjectData, err := util.GetMongoProjectData(localProjectData, buildInfo.BuildParams.City, userLogs, appData.ProjectsCollection)
	if err != nil {
		return errors.New(StartBuildsError + ": " + err.Error())
	}
	err = util.UpdateMongoProjectData(&mongoProjectData, localProjectData, buildInfo.BuildParams.City, userLogs, appData.ProjectsCollection)
	if err != nil {
		return errors.New(StartBuildsError + ": cannot get workgroups data: " + err.Error())
	}

	err = buildLoop(buildInfo, mongoProjectData.MongoWorkgroupsData[buildInfo.BuildParams.City], userLogs, appData.BuildLogsCollection)
	if err != nil {
		return err
	}
	return mongoUtil.UpdateProject(mongoProjectData.Name, mongoProjectData.Module,
		bson.M{"$set": bson.M{"mongoworkgroupsdata": mongoProjectData.MongoWorkgroupsData}}, appData.ProjectsCollection)
}
