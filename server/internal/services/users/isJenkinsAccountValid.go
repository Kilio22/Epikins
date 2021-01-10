package users

import (
	"context"
	"log"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func IsJenkinsAccountValid(newUser internal.User, credentialsCollection *mongo.Collection) internal.MyError {
	res := credentialsCollection.FindOne(context.TODO(), bson.M{"login": newUser.JenkinsLogin})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return util.GetMyError("bad jenkins account id", nil, http.StatusBadRequest)
		}
		log.Println(res.Err())
		return util.GetMyError("cannot insert given user", res.Err(), http.StatusInternalServerError)
	}
	return internal.MyError{}
}
