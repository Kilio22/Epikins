package projectsService

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/mongo"
)

func getResponseFromProjectList(projectList []libJenkins.Project, collection *mongo.Collection) ([]ProjectResponse, internal.MyError) {
	var response []ProjectResponse

	for _, project := range projectList {
		projectData, err := util.GetMongoProjectData(project, []libJenkins.Job{}, collection)
		if err != nil {
			return nil, internal.MyError{
				Err:        errors.New("cannot get projects data: " + err.Error()),
				StatusCode: http.StatusInternalServerError,
			}
		}
		response = append(response, ProjectResponse{
			Job:        project.Job,
			Module:     project.Module,
			BuildLimit: projectData.BuildLimit,
		})
	}
	return response, internal.MyError{}
}
