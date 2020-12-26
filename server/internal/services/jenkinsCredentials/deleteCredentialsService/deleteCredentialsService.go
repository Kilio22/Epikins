package deleteCredentialsService

import (
	"context"
	"epikins-api/internal"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func DeleteCredentialsService(login string, appData *internal.AppData) (myError internal.MyError) {
	res := appData.JenkinsCredentialsCollection.FindOneAndDelete(context.TODO(), bson.M{"login": login})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			myError = internal.MyError{
				Err:        errors.New("cannot delete given credentials: no credentials found with given username"),
				StatusCode: http.StatusBadRequest,
			}
		} else {
			log.Println(res.Err())
			myError = internal.MyError{
				Err:        errors.New("cannot delete given credentials: " + res.Err().Error()),
				StatusCode: http.StatusBadRequest,
			}
		}
	}
	return
}
