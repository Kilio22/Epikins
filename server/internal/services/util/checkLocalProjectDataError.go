package util

import (
	"errors"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util/mongoUtil"
	"go.mongodb.org/mongo-driver/mongo"
)

func CheckLocalProjectDataError(
	myError internal.MyError, projectName string, module string, projectCollection *mongo.Collection) internal.MyError {
	if myError.Message == "" {
		return internal.MyError{}
	}
	if myError.Status == http.StatusBadRequest {
		_ = mongoUtil.DeleteMongoProjectData(projectName, module, projectCollection)
	}
	return GetMyError("cannot get project information", errors.New(myError.Message), myError.Status)
}
