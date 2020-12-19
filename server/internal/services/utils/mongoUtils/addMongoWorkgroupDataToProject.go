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
	newMongoworkgroupData := getNewMongoworkgroupData(job)
	err := UpdateProject(project, bson.M{"$push": bson.M{"mongoworkgroupsdata": newMongoworkgroupData}}, collection)
	if err != nil {
		return internal.MongoWorkgroupData{}, errors.New("cannot add workgroup data in DB: " + err.Error())
	}
	return newMongoworkgroupData, nil
}
