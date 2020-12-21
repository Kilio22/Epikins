package getCredentialsService

import (
	"context"
	"epikins-api/internal"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

type JenkinsCredentialsPublicData struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func getPublicData(cursor *mongo.Cursor) ([]JenkinsCredentialsPublicData, internal.MyError) {
	var credentialsList []internal.JenkinsCredentials
	err := cursor.All(context.TODO(), &credentialsList)
	if err != nil {
		log.Println(err)
		return nil, internal.MyError{
			Err:        errors.New("cannot get credentials: " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}

	var publicData []JenkinsCredentialsPublicData
	for _, credentials := range credentialsList {
		publicData = append(publicData, JenkinsCredentialsPublicData{
			Id:       credentials.Id.Hex(),
			Username: credentials.Username,
		})
	}
	return publicData, internal.MyError{}
}

func GetCredentialsService(appData *internal.AppData) ([]JenkinsCredentialsPublicData, internal.MyError) {
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
