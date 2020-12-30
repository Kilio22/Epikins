package util

import (
	"errors"

	"epikins-api/pkg/libJenkins"
)

func GetAskedProject(projectList []libJenkins.Project, projectName string) (libJenkins.Project, error) {
	for idx, project := range projectList {
		if project.Job.Name == projectName {
			return projectList[idx], nil
		}
	}
	return libJenkins.Project{}, errors.New("cannot find project with name \"" + projectName + "\"")
}
