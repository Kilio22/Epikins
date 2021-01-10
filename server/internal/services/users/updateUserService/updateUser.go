package updateUserService

import (
	"context"
	"errors"
	"log"
	"net/http"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func updateUser(user internal.User, usersCollection *mongo.Collection) internal.MyError {
	_, err := usersCollection.ReplaceOne(context.TODO(), bson.M{"email": user.Email}, user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return util.GetMyError(UpdateUserError, errors.New("cannot find any user with given email"), http.StatusBadRequest)
		}
		log.Println(err)
		return util.GetMyError(UpdateUserError, err, http.StatusInternalServerError)
	}
	return internal.MyError{}
}
