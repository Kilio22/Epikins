package updateUserService

import (
	"errors"

	"epikins-api/internal"
	"epikins-api/internal/services/users"
	"epikins-api/internal/services/util"
)

const UpdateUserError = "cannot update user"

func UpdateUserService(user internal.User, appData *internal.AppData) internal.MyError {
	myError := users.IsJenkinsAccountValid(user, appData.JenkinsCredentialsCollection)
	if myError.Message != "" {
		return util.GetMyError(UpdateUserError, errors.New(myError.Message), myError.Status)
	}
	return updateUser(user, appData.UsersCollection)
}
