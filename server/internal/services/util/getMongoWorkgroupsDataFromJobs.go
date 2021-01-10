package util

import (
	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func GetMongoWorkgroupsDataFromJobs(jobs []libJenkins.Job) []internal.MongoWorkgroupData {
	var mongoWorkgroupsData []internal.MongoWorkgroupData

	for _, job := range jobs {
		mongoWorkgroupsData = append(mongoWorkgroupsData, GetNewMongoWorkgroupData(job, config.DefaultBuildNb))
	}
	return mongoWorkgroupsData
}
