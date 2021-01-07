package util

import (
	"errors"
	"time"

	"epikins-api/internal"
	"epikins-api/internal/services/util/mongoUtil"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getNewMongoWorkgroupData(job libJenkins.Job, buildLimit int) internal.MongoWorkgroupData {
	return internal.MongoWorkgroupData{
		Url:             job.Url,
		Name:            job.Name,
		RemainingBuilds: buildLimit,
		LastBuildReset:  mongoUtil.GetLastMondayDate(),
	}
}

func addMongoWorkgroupData(
	job libJenkins.Job, newWorkgroupsData []internal.MongoWorkgroupData, mongoProjectData internal.MongoProjectData,
) []internal.MongoWorkgroupData {
	if mongoWorkgroupData, ok := HasMongoWorkgroupData(job.Name, mongoProjectData.MongoWorkgroupsData); ok {
		if shouldResetWorkgroupRemainingBuilds(mongoWorkgroupData) {
			mongoWorkgroupData.RemainingBuilds = mongoProjectData.BuildLimit
		}
		newWorkgroupsData = append(newWorkgroupsData, mongoWorkgroupData)
	} else {
		newWorkgroupsData = append(newWorkgroupsData, getNewMongoWorkgroupData(job, mongoProjectData.BuildLimit))
	}
	return newWorkgroupsData
}

func updateMongoWorkgroupsData(
	mongoProjectData *internal.MongoProjectData, jobs []libJenkins.Job, projectCollection *mongo.Collection,
) error {
	var newMongoWorkgroupsData []internal.MongoWorkgroupData

	for _, job := range jobs {
		newMongoWorkgroupsData = addMongoWorkgroupData(job, newMongoWorkgroupsData, *mongoProjectData)
	}
	mongoProjectData.MongoWorkgroupsData = newMongoWorkgroupsData
	mongoProjectData.LastUpdate = time.Now().Unix()
	return mongoUtil.UpdateProject(mongoProjectData.Name,
		bson.M{
			"$set": bson.M{
				"mongoworkgroupsdata": mongoProjectData.MongoWorkgroupsData, "lastupdate": mongoProjectData.LastUpdate,
			},
		},
		projectCollection,
	)
}

// TODO: faire une requête à jenkins que si le temps est dépassé
func UpdateMongoProjectData(
	mongoProjectData *internal.MongoProjectData, jobs []libJenkins.Job, projectCollection *mongo.Collection,
) error {
	if time.Since(time.Unix(mongoProjectData.LastUpdate, 0)).Hours() < float64(24) && len(jobs) == len(mongoProjectData.MongoWorkgroupsData) {
		err := resetWorkgroupsRemainingBuilds(mongoProjectData, projectCollection)
		if err != nil {
			return errors.New("cannot update jobs remaining builds: " + err.Error())
		}
		return nil
	}
	return updateMongoWorkgroupsData(mongoProjectData, jobs, projectCollection)
}
