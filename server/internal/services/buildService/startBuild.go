package buildService

import (
	"errors"

	"epikins-api/pkg/libJenkins"
)

func startBuild(jobBuildData *GroupBuildData, buildParams BuildParams, userLogs libJenkins.Logs) (bool, error) {
	if !buildParams.FuMode {
		ok := canBuild(*jobBuildData)
		if !ok {
			return true, nil
		}
	}

	err := libJenkins.BuildJob(jobBuildData.groupJob.Url, buildParams.Visibility, userLogs)
	if err != nil {
		return false, errors.New("cannot start build: " + err.Error())
	}

	if !buildParams.FuMode {
		jobBuildData.mongoGroupData.RemainingBuilds--
	}
	return true, nil
}
