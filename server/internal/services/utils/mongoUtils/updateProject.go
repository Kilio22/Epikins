package mongoUtils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func UpdateProject(projectName string, fieldsToUpdate bson.M, collection *mongo.Collection) error {
	_, err := collection.UpdateOne(context.TODO(), bson.M{"name": projectName}, fieldsToUpdate)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
