package users

import (
	"context"
	"epikins-api/internal"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func IsJenkinsAccountValid(newUser internal.User, credentialsCollection *mongo.Collection) internal.MyError {
	res := credentialsCollection.FindOne(context.TODO(), bson.M{"login": newUser.JenkinsLogin})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return internal.MyError{
				Err:        errors.New("bad jenkins account id"),
				StatusCode: http.StatusBadRequest,
			}
		}
		log.Println(res.Err())
		return internal.MyError{
			Err:        errors.New("cannot insert given user: " + res.Err().Error()),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return internal.MyError{}
}
