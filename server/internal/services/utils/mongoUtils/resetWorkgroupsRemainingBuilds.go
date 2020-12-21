package mongoUtils

import (
	"epikins-api/config"
	"epikins-api/internal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func resetWorkgroupsRemainingBuilds(projectData *internal.MongoProjectData, collection *mongo.Collection) error {
	for idx := range projectData.MongoWorkgroupsData {
		projectData.MongoWorkgroupsData[idx].RemainingBuilds = config.DefaultBuildNb
	}
	projectData.LastUpdate = getLastMondayDate()
	return UpdateProject(projectData.Name, bson.M{
		"$set": bson.M{"mongoworkgroupsdata": projectData.MongoWorkgroupsData, "lastupdate": projectData.LastUpdate}},
		collection)
}
