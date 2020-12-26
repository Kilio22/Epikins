package addUserService

import (
	"context"
	"epikins-api/internal"
	"epikins-api/internal/services/users"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func canInsertUser(newUser internal.User, usersCollection *mongo.Collection, credentialsCollection *mongo.Collection) internal.MyError {
	res := usersCollection.FindOne(context.TODO(), bson.M{"email": newUser.Email})
	if res.Err() != nil && res.Err() != mongo.ErrNoDocuments {
		log.Println(res.Err())
		return internal.MyError{
			Err:        errors.New("cannot insert given user: " + res.Err().Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	if res.Err() == nil {
		return internal.MyError{
			Err:        errors.New("cannot insert given user: one with the same email already exists"),
			StatusCode: http.StatusConflict,
		}
	}

	myError := users.IsJenkinsAccountValid(newUser, credentialsCollection)
	if myError.Err != nil {
		return internal.MyError{
			Err:        errors.New("cannot insert given user: " + myError.Err.Error()),
			StatusCode: myError.StatusCode,
		}
	}
	return internal.MyError{}
}
