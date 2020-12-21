package addUserService

import (
	"context"
	"epikins-api/internal"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func insertUser(newUser internal.User, collection *mongo.Collection) internal.MyError {
	_, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Println(err)
		return internal.MyError{
			Err:        errors.New("cannot insert given user: " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return internal.MyError{}
}
