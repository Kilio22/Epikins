package util

import (
	"time"

	"epikins-api/internal"
	"epikins-api/internal/services/util/mongoUtil"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func shouldResetWorkgroupRemainingBuilds(workgroupData internal.MongoWorkgroupData) bool {
	return time.Since(time.Unix(workgroupData.LastBuildReset, 0)).Hours() >= float64(24*7)
}

func resetWorkgroupsRemainingBuilds(projectData *internal.MongoProjectData, city string, collection *mongo.Collection) error {
	if len(projectData.CitiesData[city].MongoWorkgroupsData) == 0 {
		return nil
	}
	if shouldResetWorkgroupRemainingBuilds(projectData.CitiesData[city].MongoWorkgroupsData[0]) {
		for idx := range projectData.CitiesData[city].MongoWorkgroupsData {
			projectData.CitiesData[city].MongoWorkgroupsData[idx].RemainingBuilds = projectData.BuildLimit
			projectData.CitiesData[city].MongoWorkgroupsData[idx].LastBuildReset = mongoUtil.GetLastMondayDate()
		}
		return mongoUtil.UpdateProject(projectData.Name, projectData.Module, bson.M{"$set": bson.M{"citiesdata": projectData.CitiesData}}, collection)
	}
	return nil
}
