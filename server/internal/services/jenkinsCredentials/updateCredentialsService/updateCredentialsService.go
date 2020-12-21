package updateCredentialsService

import (
	"context"
	"epikins-api/internal"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func UpdateCredentialsService(credentials internal.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	res := appData.JenkinsCredentialsCollection.FindOneAndUpdate(context.TODO(), bson.M{"username": credentials.Username}, credentials)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return internal.MyError{
				Err:        errors.New("cannot update given credentials: resource with given username does not exists"),
				StatusCode: http.StatusBadRequest,
			}
		}
		log.Println(res.Err())
		return internal.MyError{
			Err:        errors.New("cannot update given credentials: " + res.Err().Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return internal.MyError{}
}
