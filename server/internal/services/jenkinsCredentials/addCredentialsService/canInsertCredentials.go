package addCredentialsService

import (
	"context"
	"errors"
	"log"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func canInsertCredentials(newCredentials libJenkins.JenkinsCredentials, collection *mongo.Collection) internal.MyError {
	res := collection.FindOne(context.TODO(), bson.M{"login": newCredentials.Login})
	if res.Err() != nil && res.Err() != mongo.ErrNoDocuments {
		log.Println(res.Err())
		return util.GetMyError(AddCredentialsError, res.Err(), http.StatusInternalServerError)
	}
	if res.Err() == nil {
		return util.GetMyError(AddCredentialsError, errors.New("one with the same username already exists"), http.StatusConflict)
	}
	return internal.MyError{}
}
