package mongoUtil

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func getNewMongoWorkgroupData(job libJenkins.Job, buildLimit int) internal.MongoWorkgroupData {
	return internal.MongoWorkgroupData{
		Name:            job.Name,
		RemainingBuilds: buildLimit,
	}
}

func AddMongoWorkgroupDataToProject(job libJenkins.Job, project string, buildLimit int, collection *mongo.Collection) (internal.MongoWorkgroupData, error) {
	newMongoWorkgroupData := getNewMongoWorkgroupData(job, buildLimit)
	err := UpdateProject(project, bson.M{"$push": bson.M{"mongoworkgroupsdata": newMongoWorkgroupData}}, collection)
	if err != nil {
		return internal.MongoWorkgroupData{}, errors.New("cannot add workgroup data in DB: " + err.Error())
	}
	return newMongoWorkgroupData, nil
}
