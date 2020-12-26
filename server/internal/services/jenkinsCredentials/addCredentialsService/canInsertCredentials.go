package addCredentialsService

import (
	"context"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func canInsertCredentials(newCredentials libJenkins.JenkinsCredentials, collection *mongo.Collection) internal.MyError {
	res := collection.FindOne(context.TODO(), bson.M{"login": newCredentials.Login})
	if res.Err() != nil && res.Err() != mongo.ErrNoDocuments {
		log.Println(res.Err())
		return internal.MyError{
			Err:        errors.New("cannot insert given credentials: " + res.Err().Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	if res.Err() == nil {
		return internal.MyError{
			Err:        errors.New("cannot insert given credentials: one with the same username already exists"),
			StatusCode: http.StatusConflict,
		}
	}
	return internal.MyError{}
}
