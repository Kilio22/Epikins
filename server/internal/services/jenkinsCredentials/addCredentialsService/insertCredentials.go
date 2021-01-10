package addCredentialsService

import (
	"context"
	"log"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertCredentials(newCredentials libJenkins.JenkinsCredentials, collection *mongo.Collection) internal.MyError {
	_, err := collection.InsertOne(context.TODO(), newCredentials)
	if err != nil {
		log.Println(err)
		return util.GetMyError(AddCredentialsError, err, http.StatusInternalServerError)
	}
	return internal.MyError{}
}
