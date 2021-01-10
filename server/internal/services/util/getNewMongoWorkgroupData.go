package util

import (
	"epikins-api/internal"
	"epikins-api/internal/services/util/mongoUtil"
	"epikins-api/pkg/libJenkins"
)

func GetNewMongoWorkgroupData(job libJenkins.Job, buildLimit int) internal.MongoWorkgroupData {
	return internal.MongoWorkgroupData{
		Url:             job.Url,
		Name:            job.Name,
		RemainingBuilds: buildLimit,
		LastBuildReset:  mongoUtil.GetLastMondayDate(),
	}
}
