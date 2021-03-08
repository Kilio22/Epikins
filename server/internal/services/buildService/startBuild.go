package buildService

import (
	"errors"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/internal/services/util/mongoUtil"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/mongo"
)

const StartBuildError = "cannot start build: "

func canBuild(mongoWorkgroupData internal.MongoWorkgroupData) bool {
	return mongoWorkgroupData.RemainingBuilds > 0
}

func startBuild(
	mongoWorkgroupData *internal.MongoWorkgroupData, buildInfo BuildInfo, userLogs libJenkins.JenkinsCredentials,
	buildLogsCollection *mongo.Collection) (
	bool, error,
) {
	if !buildInfo.BuildParams.Fu {
		ok := canBuild(*mongoWorkgroupData)
		if !ok {
			return true, nil
		}
	}

	err := libJenkins.BuildJob(mongoWorkgroupData.Url, buildInfo.BuildParams.Visibility, userLogs)
	if err != nil {
		return false, errors.New(StartBuildError + err.Error())
	}

	_, err = mongoUtil.AddBuildLog(util.GetNewBuildLog(buildInfo.BuildParams.Module, buildInfo.Starter, mongoWorkgroupData.Name, buildInfo.BuildParams.Project), buildLogsCollection)
	if err != nil {
		return true, errors.New(StartBuildError + err.Error())
	}

	if !buildInfo.BuildParams.Fu {
		mongoWorkgroupData.RemainingBuilds--
	}
	return true, nil
}
