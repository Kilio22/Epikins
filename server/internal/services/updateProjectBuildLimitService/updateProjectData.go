package updateProjectBuildLimitService

import (
	"context"

	"epikins-api/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func updateValues(newLimit NewLimit, mongoProjectData *internal.MongoProjectData) {
	oldLimit := mongoProjectData.BuildLimit
	mongoProjectData.BuildLimit = newLimit.BuildLimit

	for key, mongoWorkgroupsData := range mongoProjectData.MongoWorkgroupsData {
		for idx, workgroupData := range mongoWorkgroupsData {
			if newLimit.BuildLimit-(oldLimit-workgroupData.RemainingBuilds) >= 0 {
				mongoProjectData.MongoWorkgroupsData[key][idx].RemainingBuilds = newLimit.BuildLimit - (oldLimit - workgroupData.RemainingBuilds)
			} else {
				mongoProjectData.MongoWorkgroupsData[key][idx].RemainingBuilds = 0
			}
		}
	}
}

func updateProjectData(newLimit NewLimit, projectName string, module string, collection *mongo.Collection) error {
	var projectData internal.MongoProjectData
	err := collection.FindOne(context.TODO(), bson.M{"name": projectName, "module": module}).Decode(&projectData)
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
