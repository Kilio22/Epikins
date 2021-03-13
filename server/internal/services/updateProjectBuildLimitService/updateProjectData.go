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

	for key, cityData := range mongoProjectData.CitiesData {
		for idx, workgroupData := range cityData.MongoWorkgroupsData {
			if newLimit.BuildLimit-(oldLimit-workgroupData.RemainingBuilds) >= 0 {
				mongoProjectData.CitiesData[key].MongoWorkgroupsData[idx].RemainingBuilds = newLimit.BuildLimit - (oldLimit - workgroupData.RemainingBuilds)
			} else {
				mongoProjectData.CitiesData[key].MongoWorkgroupsData[idx].RemainingBuilds = 0
			}
		}
	}
}

func updateProjectData(newLimit NewLimit, projectName string, module string, collection *mongo.Collection) error {
	var mongoProjectData internal.MongoProjectData
	err := collection.FindOne(context.TODO(), bson.M{"name": projectName, "module": module}).Decode(&mongoProjectData)
	if err != nil {
		return err
	}
	if mongoProjectData.BuildLimit == newLimit.BuildLimit {
		return nil
	}
	updateValues(newLimit, &mongoProjectData)
	_, err = collection.ReplaceOne(context.TODO(), bson.M{"name": projectName, "module": module}, mongoProjectData)
	return err
}
