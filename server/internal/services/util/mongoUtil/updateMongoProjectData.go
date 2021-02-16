package mongoUtil

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateProject(projectName string, module string, fieldsToUpdate bson.M, collection *mongo.Collection) error {
	res := collection.FindOneAndUpdate(context.TODO(), bson.M{"name": projectName, "module": module}, fieldsToUpdate)
	if res.Err() != nil {
		log.Println(res.Err())
		return res.Err()
	}
	return nil
}
