package mongoUtils

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func getNewMongoworkgroupData(job libJenkins.Job) internal.MongoWorkgroupData {
	return internal.MongoWorkgroupData{
		Name:            job.Name,
		RemainingBuilds: config.DefaultBuildNb,
	}
}

func AddMongoWorkgroupDataToProject(job libJenkins.Job, project string, collection *mongo.Collection) (internal.MongoWorkgroupData, error) {
	newMongoWorkgroupData := getNewMongoworkgroupData(job)
	err := UpdateProject(project, bson.M{"$push": bson.M{"mongoworkgroupsdata": newMongoWorkgroupData}}, collection)
	if err != nil {
		return internal.MongoWorkgroupData{}, errors.New("cannot add workgroup data in DB: " + err.Error())
	}
	return newMongoWorkgroupData, nil
}
