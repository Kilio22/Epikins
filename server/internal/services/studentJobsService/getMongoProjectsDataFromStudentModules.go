package studentJobsService

import (
	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

func getModulesNameFromModules(modules []Module) []string {
	var moduleNameList []string

	for _, module := range modules {
		moduleNameList = append(moduleNameList, module.CodeModule)
	}
	return moduleNameList
}

func getMongoProjectsDataByModule(
	moduleName string, city string, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) (
	[]internal.MongoProjectData, error,
) {
	var moduleMongoProjectsData []internal.MongoProjectData

	for _, projectData := range appData.ProjectsData[userLogs.Login].ProjectList {
		if projectData.Module == moduleName {
			mongoProjectData, err := util.GetMongoProjectData(projectData, city, userLogs, appData.ProjectsCollection)
			if err != nil {
				return []internal.MongoProjectData{}, err
			}
			moduleMongoProjectsData = append(moduleMongoProjectsData, mongoProjectData)
		}
	}
	return moduleMongoProjectsData, nil
}

func getMongoProjectsDataFromStudentModules(
	modules []Module, city string, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) ([]internal.MongoProjectData, error) {
	if err := util.CheckLocalProjectsData(userLogs, false, appData); err != nil {
		return []internal.MongoProjectData{}, err
	}

	modulesName := getModulesNameFromModules(modules)
	var mongoProjectsData []internal.MongoProjectData
	for _, moduleName := range modulesName {
		moduleMongoProjectData, err := getMongoProjectsDataByModule(moduleName, city, userLogs, appData)
		if err != nil {
			return []internal.MongoProjectData{}, err
		}
		mongoProjectsData = append(mongoProjectsData, moduleMongoProjectData...)
	}
	return mongoProjectsData, nil
}
