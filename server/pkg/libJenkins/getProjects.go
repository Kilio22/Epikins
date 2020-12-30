package libJenkins

import (
	"errors"
)

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
			fullProjectList = append(fullProjectList, Project{
				Job:    project,
				Module: module.Name,
			})
		}
	}
	return fullProjectList, nil
}
