package addUserService

import (
	"context"
	"log"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/mongo"
)

func insertUser(newUser internal.User, collection *mongo.Collection) internal.MyError {
	_, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Println(err)
		return util.GetMyError("cannot insert given user", err, http.StatusInternalServerError)
	}
	return internal.MyError{}
}
