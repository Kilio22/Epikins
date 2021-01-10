package deleteCredentialsService

import (
	"context"
	"errors"
	"log"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const DeleteCredentialsError = "cannot delete given credentials"

func DeleteCredentialsService(login string, appData *internal.AppData) internal.MyError {
	res := appData.JenkinsCredentialsCollection.FindOneAndDelete(context.TODO(), bson.M{"login": login})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return util.GetMyError(DeleteCredentialsError, errors.New("no credentials found with given username"), http.StatusBadRequest)
		}
		log.Println(res.Err())
		return util.GetMyError(DeleteCredentialsError, res.Err(), http.StatusInternalServerError)
	}
	return internal.MyError{}
}
