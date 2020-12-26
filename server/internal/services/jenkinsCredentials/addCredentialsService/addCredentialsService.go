package addCredentialsService

import (
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func AddCredentialsService(newCredentials libJenkins.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	myError := canInsertCredentials(newCredentials, appData.JenkinsCredentialsCollection)
	if myError.Err != nil {
		return myError
	}
	return insertCredentials(newCredentials, appData.JenkinsCredentialsCollection)
}
