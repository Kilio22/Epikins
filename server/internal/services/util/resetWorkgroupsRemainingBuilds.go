package util

import (
	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/services/util/mongoUtil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func shouldResetWorkgroupRemainingBuilds(workgroupData internal.MongoWorkgroupData) bool {
	return time.Since(time.Unix(workgroupData.LastBuildReset, 0)).Hours() >= float64(24*7)
}

func resetWorkgroupsRemainingBuilds(projectData *internal.MongoProjectData, collection *mongo.Collection) error {
	if len(projectData.MongoWorkgroupsData) == 0 {
		return nil
	}
	if shouldResetWorkgroupRemainingBuilds(projectData.MongoWorkgroupsData[0]) {
		for idx := range projectData.MongoWorkgroupsData {
			projectData.MongoWorkgroupsData[idx].RemainingBuilds = config.DefaultBuildNb
		}
		return mongoUtil.UpdateProject(projectData.Name, bson.M{"$set": bson.M{"mongoworkgroupsdata": projectData.MongoWorkgroupsData}}, collection)
	}
	return nil
}
