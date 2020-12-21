package updateUserService

import (
	"context"
	"epikins-api/internal"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func UpdateUserService(user internal.User, appData *internal.AppData) internal.MyError {
	res := appData.UsersCollection.FindOneAndUpdate(context.TODO(), bson.M{"email": user.Email}, user)
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return internal.MyError{
				Err:        errors.New("cannot update given user: no resource with given email exists"),
				StatusCode: http.StatusBadRequest,
			}
		}
		log.Println(res.Err())
		return internal.MyError{
			Err:        errors.New("cannot update given user: " + res.Err().Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return internal.MyError{}
}
