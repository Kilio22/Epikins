package addCredentialsService

import (
	"epikins-api/internal"
)

func AddCredentialsService(newCredentials internal.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	myError := canInsertCredentials(newCredentials, appData.JenkinsCredentialsCollection)
	if myError.Err != nil {
		return myError
	}
	return insertCredentials(newCredentials, appData.JenkinsCredentialsCollection)
}
