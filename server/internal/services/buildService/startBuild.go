package buildService

import (
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func canBuild(mongoWorkgroupData internal.MongoWorkgroupData) bool {
	return mongoWorkgroupData.RemainingBuilds > 0
}

func startBuild(mongoWorkgroupData *internal.MongoWorkgroupData, buildParams BuildParams, userLogs libJenkins.JenkinsCredentials) (
	bool, error,
) {
	if !buildParams.FuMode {
		ok := canBuild(*mongoWorkgroupData)
		if !ok {
			return true, nil
		}
	}

	/*err := libJenkins.BuildJob(mongoWorkgroupData.Url, buildParams.Visibility, userLogs)
	if err != nil {
		return false, errors.New("cannot start build: " + err.Error())
	}
	*/
	if !buildParams.FuMode {
		mongoWorkgroupData.RemainingBuilds--
	}
	return true, nil
}
