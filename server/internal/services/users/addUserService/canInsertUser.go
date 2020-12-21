package addUserService

import (
	"context"
	"epikins-api/internal"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func isJenkinsAccountValid(newUser internal.User, credentialsCollection *mongo.Collection) internal.MyError {
	objectId, err := primitive.ObjectIDFromHex(newUser.JenkinsAccountId)
	if err != nil {
		return internal.MyError{
			Err:        errors.New("cannot insert given user: bad jenkins account id"),
			StatusCode: http.StatusBadRequest,
		}
	}

	res := credentialsCollection.FindOne(context.TODO(), bson.M{"_id": objectId})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return internal.MyError{
				Err:        errors.New("cannot insert given user: bad jenkins account id"),
				StatusCode: http.StatusBadRequest,
			}
		}
		return internal.MyError{
			Err:        errors.New("cannot insert given user: " + res.Err().Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return internal.MyError{}
}

func canInsertUser(newUser internal.User, usersCollection *mongo.Collection, credentialsCollection *mongo.Collection) internal.MyError {
	res := usersCollection.FindOne(context.TODO(), bson.M{"email": newUser.Email})
	println(res)
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
	return isJenkinsAccountValid(newUser, credentialsCollection)
}
