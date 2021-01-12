package projectsService

import (
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/mongo"
)

func getResponseFromProjectList(
	projectList []libJenkins.Project, userLogs libJenkins.JenkinsCredentials, collection *mongo.Collection) (
	[]ProjectResponse, internal.MyError,
) {
	var response []ProjectResponse

	for _, project := range projectList {
		projectData, err := util.GetMongoProjectData(project, "", userLogs, collection)
		if err != nil {
			return nil, util.GetMyError(ProjectsError, err, http.StatusInternalServerError)
		}
		response = append(response, ProjectResponse{
			Cities:     project.Cities,
			Job:        project.Job,
			Module:     project.Module,
			BuildLimit: projectData.BuildLimit,
		})
	}
	return response, internal.MyError{}
}
