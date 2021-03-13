package util

import (
	"time"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func GetCitiesDataMapFromJobs(jobs []libJenkins.Job, city string) map[string]internal.CityData {
	var mongoWorkgroupsData []internal.MongoWorkgroupData

	if city == "" {
		return map[string]internal.CityData{}
	}
	for _, job := range jobs {
		mongoWorkgroupsData = append(mongoWorkgroupsData, GetNewMongoWorkgroupData(job, config.DefaultBuildNb))
	}
	return map[string]internal.CityData{
		city: {
			LastUpdate:          time.Now().Unix(),
			MongoWorkgroupsData: mongoWorkgroupsData,
		},
	}
}
