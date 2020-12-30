package updateProjectBuildLimitService

import (
	"epikins-api/internal"
	"epikins-api/internal/services/utils/mongoUtils"
	"epikins-api/pkg/libJenkins"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type NewLimit struct {
	BuildLimit int `json:"buildLimit" validate:"gte=0"`
}

func checkError(err error, shouldAddProject bool, projectName string, jenkinsCredentials libJenkins.JenkinsCredentials, appData *internal.AppData) (bool, internal.MyError) {
	if err == nil {
		return false, internal.MyError{}
	}
	if err == mongo.ErrNoDocuments && shouldAddProject {
		ok, myError := canAddProject(projectName, jenkinsCredentials, appData)
		if ok {
			return true, internal.MyError{}
		}
		if !ok && myError.Err == nil {
			return false, internal.MyError{
				Err:        errors.New("cannot update projectName build limit: no project with name \"" + projectName + "\" were found"),
				StatusCode: http.StatusBadRequest,
			}
		}
		return false, internal.MyError{
			Err:        errors.New("cannot update projectName build limit: " + myError.Err.Error()),
			StatusCode: myError.StatusCode,
		}
	}
	return false, internal.MyError{
		Err:        errors.New("cannot update projectName build limit: " + err.Error()),
		StatusCode: http.StatusInternalServerError,
	}
}

func UpdateProjectBuildLimitService(newLimit NewLimit, projectName string, jenkinsCredentials libJenkins.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	err := updateProjectData(newLimit, projectName, appData.ProjectsCollection)
	shouldRetry, myError := checkError(err, true, projectName, jenkinsCredentials, appData)

	if shouldRetry && myError.Err == nil {
		_, err := mongoUtils.AddMongoProjectData(projectName, []libJenkins.Job{}, appData.ProjectsCollection)
		if err != nil {
			return internal.MyError{
				Err:        errors.New("cannot update build limit: something went wrong when trying to add projectName in DB: " + err.Error()),
				StatusCode: http.StatusInternalServerError,
			}
		}
		err = updateProjectData(newLimit, projectName, appData.ProjectsCollection)
		_, myError = checkError(err, false, projectName, jenkinsCredentials, appData)
		return myError
	}
	return myError
}
