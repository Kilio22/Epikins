package deleteUserService

import (
	"context"
	"epikins-api/internal"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func DeleteUserService(username string, appData *internal.AppData) (myError internal.MyError) {
	res := appData.UsersCollection.FindOneAndDelete(context.TODO(), bson.M{"email": username})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			myError = internal.MyError{
				Err:        errors.New("cannot delete given user: no user found with given email"),
				StatusCode: http.StatusBadRequest,
			}
		} else {
			log.Println(res.Err())
			myError = internal.MyError{
				Err:        errors.New("cannot delete given user: " + res.Err().Error()),
				StatusCode: http.StatusBadRequest,
			}
		}
	}
	return
}
