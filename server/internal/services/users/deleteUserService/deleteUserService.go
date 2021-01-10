package deleteUserService

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

const DeleteUserError = "cannot delete user"

func DeleteUserService(username string, appData *internal.AppData) internal.MyError {
	res := appData.UsersCollection.FindOneAndDelete(context.TODO(), bson.M{"email": username})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return util.GetMyError(DeleteUserError, errors.New("no user found with given email"), http.StatusBadRequest)
		} else {
			log.Println(res.Err())
			return util.GetMyError(DeleteUserError, res.Err(), http.StatusInternalServerError)
		}
	}
	return internal.MyError{}
}
