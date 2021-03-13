package util

import (
	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func GetNewMongoProjectData(
	project libJenkins.Project, citiesData map[string]internal.CityData) internal.MongoProjectData {
	return internal.MongoProjectData{
		BuildLimit: config.DefaultBuildNb,
		Module:     project.Module,
		CitiesData: citiesData,
		Name:       project.Job.Name,
	}
}
