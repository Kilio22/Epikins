package addUserService

import "epikins-api/internal"

const AddUserError = "cannot insert user"

func AddUserService(newUser internal.User, appData *internal.AppData) internal.MyError {
	myError := canInsertUser(newUser, appData.UsersCollection, appData.JenkinsCredentialsCollection)
	if myError.Message != "" {
		return myError
	}
	return insertUser(newUser, appData.UsersCollection)
}
