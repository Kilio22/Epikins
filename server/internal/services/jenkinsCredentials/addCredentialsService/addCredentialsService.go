package addCredentialsService

import (
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

const AddCredentialsError = "cannot insert given credentials"

func AddCredentialsService(newCredentials libJenkins.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	myError := canInsertCredentials(newCredentials, appData.JenkinsCredentialsCollection)
	if myError.Message != "" {
		return myError
	}
	return insertCredentials(newCredentials, appData.JenkinsCredentialsCollection)
}
