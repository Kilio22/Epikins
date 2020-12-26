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

func updateUser(user internal.User, usersCollection *mongo.Collection) internal.MyError {
	_, err := usersCollection.ReplaceOne(context.TODO(), bson.M{"email": user.Email}, user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return internal.MyError{
				Err:        errors.New("cannot update given user: no resource with given email exists"),
				StatusCode: http.StatusBadRequest,
			}
		}
		log.Println(err)
		return internal.MyError{
			Err:        errors.New("cannot update given user: " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return internal.MyError{}
}
