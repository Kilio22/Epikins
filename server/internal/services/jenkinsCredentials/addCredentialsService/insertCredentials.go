package addCredentialsService

import (
	"context"
	"epikins-api/internal"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func insertCredentials(newCredentials internal.JenkinsCredentials, collection *mongo.Collection) internal.MyError {
	_, err := collection.InsertOne(context.TODO(), newCredentials)
	if err != nil {
		log.Println(err)
		return internal.MyError{
			Err:        errors.New("cannot insert given credentials: " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return internal.MyError{}
}
