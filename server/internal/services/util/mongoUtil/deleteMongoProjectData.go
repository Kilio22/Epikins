package mongoUtil

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func DeleteMongoProjectData(projectName string, module string, projectCollection *mongo.Collection) error {
	_, err := projectCollection.DeleteOne(context.TODO(), bson.M{"name": projectName, "module": module})
	return err
}
