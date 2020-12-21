package addUserService

import "epikins-api/internal"

func AddUserService(newUser internal.User, appData *internal.AppData) internal.MyError {
	myError := canInsertUser(newUser, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if myError.Err != nil {
		return myError
	}
	return insertUser(newUser, appData.UsersCollection)
}
