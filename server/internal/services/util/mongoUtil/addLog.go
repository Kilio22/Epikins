package mongoUtil

import (
	"context"
	"errors"
	"log"

	"epikins-api/internal"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddBuildLog(newBuildLog internal.BuildLogElem, collection *mongo.Collection) (
	internal.BuildLogElem, error,
) {
	_, err := collection.InsertOne(context.TODO(), newBuildLog)
	if err != nil {
		log.Println(err)
		return internal.BuildLogElem{}, errors.New("cannot add build log in db: " + err.Error())
	}
	return newBuildLog, nil
}
