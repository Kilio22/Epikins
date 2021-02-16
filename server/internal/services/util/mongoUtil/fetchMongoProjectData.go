package mongoUtil

import (
	"context"

	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FetchMongoProjectData(project libJenkins.Project, collection *mongo.Collection) (internal.MongoProjectData, error) {
	var projectData internal.MongoProjectData
	err := collection.FindOne(
		context.TODO(), bson.M{"name": project.Job.Name, "module": project.Module}).Decode(&projectData)
	return projectData, err
}
