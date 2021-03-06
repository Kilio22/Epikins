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
	buildLogCollection *mongo.Collection) error {
	for _, jobName := range buildInfo.BuildParams.Jobs {
		for idx := range groupsBuildData {
			if groupsBuildData[idx].Name != jobName {
				continue
			}
			shouldBreak, err := startBuild(&groupsBuildData[idx], buildInfo, userLogs, buildLogCollection)
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
	err = util.UpdateMongoProjectData(&mongoProjectData, localProjectData, buildInfo.BuildParams.City, false, userLogs, appData.ProjectsCollection)
	if err != nil {
		return errors.New(StartBuildsError + ": cannot get workgroups data: " + err.Error())
	}

	err = buildLoop(buildInfo, mongoProjectData.CitiesData[buildInfo.BuildParams.City].MongoWorkgroupsData, userLogs, appData.BuildLogCollection)
	if err != nil {
		return err
	}
	return mongoUtil.UpdateProject(mongoProjectData.Name, mongoProjectData.Module,
		bson.M{"$set": bson.M{"citiesdata": mongoProjectData.CitiesData}}, appData.ProjectsCollection)
}
