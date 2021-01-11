package util

import (
	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func GetMongoWorkgroupsDataFromJobs(jobs []libJenkins.Job, city string) map[string][]internal.MongoWorkgroupData {
	var mongoWorkgroupsData []internal.MongoWorkgroupData

	if city == "" {
		return map[string][]internal.MongoWorkgroupData{}
	}
	for _, job := range jobs {
		mongoWorkgroupsData = append(mongoWorkgroupsData, GetNewMongoWorkgroupData(job, config.DefaultBuildNb))
	}
	return map[string][]internal.MongoWorkgroupData{
		city: mongoWorkgroupsData,
	}
}
