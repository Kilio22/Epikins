package mongoUtil

import (
	"context"
	"errors"
	"log"

	"epikins-api/internal"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddBuildLog(newBuildLog internal.BuildLog, collection *mongo.Collection) (
	internal.BuildLog, error,
) {
	_, err := collection.InsertOne(context.TODO(), newBuildLog)
	if err != nil {
		log.Println(err)
		return internal.BuildLog{}, errors.New("cannot add build log in db: " + err.Error())
	}
	return newBuildLog, nil
}
