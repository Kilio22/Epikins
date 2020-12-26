package libJenkins

import (
	"errors"
)

func GetProjects(userLogs JenkinsCredentials) ([]Job, error) {
	moduleList, err := GetJobsByURL(JenkinsBaseURL, userLogs)
	if err != nil {
		return []Job{}, errors.New("cannot get projects: something went wrong when reaching module list: " + err.Error())
	}

	var fullProjectList []Job
	for _, module := range moduleList {
		projectList, err := GetJobsByURL(module.Url, userLogs)
		if err != nil {
			return []Job{}, errors.New("cannot get projects: something went wrong when reaching project list for module \"" + module.Name + "\": " + err.Error())
		}
		fullProjectList = append(fullProjectList, projectList...)
	}
	return fullProjectList, nil
}
