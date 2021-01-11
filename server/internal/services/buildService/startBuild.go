package buildService

import (
	"errors"

	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

const StartBuildError = "cannot start build: "

func canBuild(mongoWorkgroupData internal.MongoWorkgroupData) bool {
	return mongoWorkgroupData.RemainingBuilds > 0
}

func startBuild(mongoWorkgroupData *internal.MongoWorkgroupData, buildParams BuildParams, userLogs libJenkins.JenkinsCredentials) (
	bool, error,
) {
	if !buildParams.Fu {
		ok := canBuild(*mongoWorkgroupData)
		if !ok {
			return true, nil
		}
	}

	err := libJenkins.BuildJob(mongoWorkgroupData.Url, buildParams.Visibility, userLogs)
	if err != nil {
		return false, errors.New(StartBuildError + err.Error())
	}
	if !buildParams.Fu {
		mongoWorkgroupData.RemainingBuilds--
	}
	return true, nil
}
