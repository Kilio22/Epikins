package libJenkins

import (
	"errors"
	"strings"
)

func hasElement(list []string, toFind string) bool {
	for _, element := range list {
		if element == toFind {
			return true
		}
	}
	return false
}

func getCitiesNameFromJobs(cities []Job) []string {
	var citiesName []string

	for _, city := range cities {
		cityName := strings.Split(city.Name, "-")
		if !hasElement(citiesName, cityName[0]) {
			citiesName = append(citiesName, cityName[0])
		}
	}
	return citiesName
}

func getCitiesFromProject(project Job, userLogs JenkinsCredentials) ([]string, error) {
	yearList, err := getYearList(project, userLogs)
	if err != nil || len(yearList) == 0 {
		return []string{}, err
	}
	sortYearList(yearList)
	cities, err := getCities(yearList[0].Url, userLogs)
	if err != nil || len(cities) == 0 {
		return []string{}, err
	}
	return getCitiesNameFromJobs(cities), nil
}

func GetProjects(userLogs JenkinsCredentials) ([]Project, error) {
	moduleList, err := GetJobsByURL(JenkinsBaseURL, userLogs)
	if err != nil {
		return []Project{}, errors.New("cannot get projects: something went wrong when reaching module list: " + err.Error())
	}

	var fullProjectList []Project
	for _, module := range moduleList {
		projectList, err := GetJobsByURL(module.Url, userLogs)
		if err != nil {
			return []Project{}, errors.New("cannot get projects: something went wrong when reaching project list for module \"" + module.Name + "\": " + err.Error())
		}
		for _, project := range projectList {
			cities, err := getCitiesFromProject(project, userLogs)
			if err != nil {
				return []Project{}, errors.New("cannot get projects: " + err.Error())
			}
			fullProjectList = append(fullProjectList, Project{
				Cities: cities,
				Job:    project,
				Module: module.Name,
			})
		}
	}
	return fullProjectList, nil
}
