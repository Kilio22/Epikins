package mongoUtils

import (
	"context"
	"epikins-api/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchMongoProjectData(project string, collection *mongo.Collection) (internal.MongoProjectData, error) {
	var projectData internal.MongoProjectData
	err := collection.FindOne(
		context.TODO(), bson.M{"name": project}).Decode(&projectData)
	return projectData, err
}
