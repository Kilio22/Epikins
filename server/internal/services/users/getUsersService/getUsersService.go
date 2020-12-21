package getUsersService

import (
	"context"
	"epikins-api/internal"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func getUsers(cursor *mongo.Cursor) ([]internal.User, internal.MyError) {
	var users []internal.User
	err := cursor.All(context.TODO(), &users)
	if err != nil {
		log.Println(err)
		return nil, internal.MyError{
			Err:        errors.New("cannot get credentials: " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return users, internal.MyError{}
}

func GetUsersService(appData *internal.AppData) ([]internal.User, internal.MyError) {
	cursor, err := appData.UsersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
		return nil, internal.MyError{
			Err:        errors.New("cannot get credentials: " + err.Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return getUsers(cursor)
}
