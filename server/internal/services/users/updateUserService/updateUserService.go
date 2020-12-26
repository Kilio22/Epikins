package updateUserService

import (
	"epikins-api/internal"
	"epikins-api/internal/services/users"
	"errors"
)

func UpdateUserService(user internal.User, appData *internal.AppData) internal.MyError {
	myError := users.IsJenkinsAccountValid(user, appData.JenkinsCredentialsCollection)
	if myError.Err != nil {
		return internal.MyError{
			Err:        errors.New("cannot update given user: " + myError.Err.Error()),
			StatusCode: myError.StatusCode,
		}
	}
	return updateUser(user, appData.UsersCollection)
}
