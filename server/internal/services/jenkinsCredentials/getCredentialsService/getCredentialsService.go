package getCredentialsService

import (
	"context"
	"log"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const GetCredentialsError = "cannot get credentials"

func getPublicData(cursor *mongo.Cursor) ([]string, internal.MyError) {
	var credentialsList []libJenkins.JenkinsCredentials
	err := cursor.All(context.TODO(), &credentialsList)
	if err != nil {
		log.Println(err)
		return nil, util.GetMyError(GetCredentialsError, err, http.StatusInternalServerError)
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
		return nil, util.GetMyError(GetCredentialsError, err, http.StatusInternalServerError)
	}
	return getPublicData(cursor)
}
