package addUserService

import (
	"context"
	"errors"
	"log"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/users"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func canInsertUser(newUser internal.User, usersCollection *mongo.Collection, credentialsCollection *mongo.Collection) internal.MyError {
	res := usersCollection.FindOne(context.TODO(), bson.M{"email": newUser.Email})
	if res.Err() != nil && res.Err() != mongo.ErrNoDocuments {
		log.Println(res.Err())
		return util.GetMyError(AddUserError, res.Err(), http.StatusInternalServerError)
	}
	if res.Err() == nil {
		return util.GetMyError(AddUserError, errors.New("one with the same email already exists"), http.StatusConflict)
	}

	myError := users.IsJenkinsAccountValid(newUser, credentialsCollection)
	if myError.Message != "" {
		return util.GetMyError(AddUserError, errors.New(myError.Message), myError.Status)
	}
	return internal.MyError{}
}
