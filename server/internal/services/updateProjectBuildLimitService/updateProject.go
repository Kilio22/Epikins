package updateProjectBuildLimitService

import (
	"context"

	"epikins-api/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func updateValues(newLimit NewLimit, projectData *internal.MongoProjectData) {
	oldLimit := projectData.BuildLimit
	projectData.BuildLimit = newLimit.BuildLimit
	for idx, workgroupData := range projectData.MongoWorkgroupsData {
		if newLimit.BuildLimit-(oldLimit-workgroupData.RemainingBuilds) >= 0 {
			projectData.MongoWorkgroupsData[idx].RemainingBuilds = newLimit.BuildLimit - (oldLimit - workgroupData.RemainingBuilds)
		} else {
			projectData.MongoWorkgroupsData[idx].RemainingBuilds = 0
		}
	}
}

func updateProjectData(newLimit NewLimit, projectName string, collection *mongo.Collection) error {
	var projectData internal.MongoProjectData
	err := collection.FindOne(context.TODO(), bson.M{"name": projectName}).Decode(&projectData)
	if err != nil {
		return err
	}
	if projectData.BuildLimit == newLimit.BuildLimit {
		return nil
	}
	updateValues(newLimit, &projectData)
	_, err = collection.ReplaceOne(context.TODO(), bson.M{"name": projectName}, projectData)
	return err
}
