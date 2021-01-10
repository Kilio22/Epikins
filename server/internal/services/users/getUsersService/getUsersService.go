package getUsersService

import (
	"context"
	"log"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const GetUsersError = "cannot get users"

func getUsers(cursor *mongo.Cursor) ([]internal.User, internal.MyError) {
	var users []internal.User
	err := cursor.All(context.TODO(), &users)
	if err != nil {
		log.Println(err)
		return nil, util.GetMyError(GetUsersError, err, http.StatusInternalServerError)
	}
	return users, internal.MyError{}
}

func GetUsersService(appData *internal.AppData) ([]internal.User, internal.MyError) {
	cursor, err := appData.UsersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Println(err)
		return nil, util.GetMyError(GetUsersError, err, http.StatusInternalServerError)
	}
	return getUsers(cursor)
}
