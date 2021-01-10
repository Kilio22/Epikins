package mongoUtil

import (
	"context"
	"errors"
	"log"

	"epikins-api/internal"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddMongoProjectData(newProjectData internal.MongoProjectData, collection *mongo.Collection) (
	internal.MongoProjectData, error,
) {
	_, err := collection.InsertOne(context.TODO(), newProjectData)
	if err != nil {
		log.Println(err)
		return internal.MongoProjectData{}, errors.New("cannot add new projectName in db: " + err.Error())
	}
	return newProjectData, nil
}
