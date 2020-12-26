package getCredentialsService

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

func getPublicData(cursor *mongo.Cursor) ([]string, internal.MyError) {
	var credentialsList []libJenkins.JenkinsCredentials
	err := cursor.All(context.TODO(), &credentialsList)
	if err != nil {
		log.Println(err)
		return nil, internal.MyError{
			Err:        errors.New("cannot get credentials: " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}

	var publicData []string
	for _, credentials := range credentialsList {
		publicData = append(publicData, credentials.Login)
	}
	return publicData, internal.MyError{}
}

func GetCredentialsService(appData *internal.AppData) ([]string, internal.MyError) {
	cursor, err := appData.JenkinsCredentialsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
		return nil, internal.MyError{
			Err:        errors.New("cannot get credentials: " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return getPublicData(cursor)
}
