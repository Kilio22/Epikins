package util

import (
	"time"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func GetNewMongoProjectData(
	project libJenkins.Project, mongoWorkgroupsData map[string][]internal.MongoWorkgroupData) internal.MongoProjectData {
	return internal.MongoProjectData{
		BuildLimit:          config.DefaultBuildNb,
		LastUpdate:          time.Now().Unix(),
		Module:              project.Module,
		MongoWorkgroupsData: mongoWorkgroupsData,
		Name:                project.Job.Name,
	}
}
