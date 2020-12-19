package buildService

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"epikins-api/internal"
	"epikins-api/internal/services/utils/mongoUtils"
)

func updateMongoWorkgroupsData(project string, jobsBuildData []GroupBuildData, collection *mongo.Collection) error {
	var updatedMongoWorkgroupsData []internal.MongoWorkgroupData
	for _, jobBuildData := range jobsBuildData {
		updatedMongoWorkgroupsData = append(updatedMongoWorkgroupsData, jobBuildData.mongoGroupData)
	}
	return mongoUtils.UpdateProject(project, bson.M{"$set": bson.M{"mongoworkgroupsdata": updatedMongoWorkgroupsData}}, collection)
}
