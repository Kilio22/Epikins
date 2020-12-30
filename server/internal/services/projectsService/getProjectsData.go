package projectsService

import (
	"epikins-api/internal"
	"epikins-api/internal/services/utils"
	"epikins-api/pkg/libJenkins"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func getProjectData(projectList []libJenkins.Project, collection *mongo.Collection) ([]ProjectResponse, internal.MyError) {
	var response []ProjectResponse

	for _, project := range projectList {
		projectData, err := utils.FetchProjectData(project.Job.Name, []libJenkins.Job{}, collection)
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
