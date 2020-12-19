package utils

import (
	"errors"

	"epikins-api/pkg/libJenkins"
)

func GetAskedProject(projectList []libJenkins.Job, projectName string) (libJenkins.Job, error) {
	for idx, project := range projectList {
		if project.Name == projectName {
			return projectList[idx], nil
		}
	}
	return libJenkins.Job{}, errors.New("cannot find project with name \"" + projectName + "\"")
}
