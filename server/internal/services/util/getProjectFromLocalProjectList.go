package util

import (
	"errors"

	"epikins-api/pkg/libJenkins"
)

func GetProjectFromLocalProjectList(projectList []libJenkins.Project, projectName string, module string) (libJenkins.Project, error) {
	for idx, project := range projectList {
		if project.Job.Name == projectName && project.Module == module {
			return projectList[idx], nil
		}
	}
	return libJenkins.Project{}, errors.New("cannot find project with name \"" + projectName + "\"")
}
