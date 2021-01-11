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

const UpdateMongoProjectDataError = "cannot update project data: "

func addMongoWorkgroupData(
	job libJenkins.Job, city string, newWorkgroupsData []internal.MongoWorkgroupData, mongoProjectData internal.MongoProjectData,
) []internal.MongoWorkgroupData {
	if mongoWorkgroupData, ok := HasMongoWorkgroupData(job.Name, mongoProjectData.MongoWorkgroupsData[city]); ok {
		if shouldResetWorkgroupRemainingBuilds(mongoWorkgroupData) {
			mongoWorkgroupData.RemainingBuilds = mongoProjectData.BuildLimit
		}
		newWorkgroupsData = append(newWorkgroupsData, mongoWorkgroupData)
	} else {
		newWorkgroupsData = append(newWorkgroupsData, GetNewMongoWorkgroupData(job, mongoProjectData.BuildLimit))
	}
	return newWorkgroupsData
}

func updateMongoWorkgroupsData(
	mongoProjectData *internal.MongoProjectData, jobs []libJenkins.Job, city string, projectCollection *mongo.Collection,
) error {
	var newMongoWorkgroupsData []internal.MongoWorkgroupData

	for _, job := range jobs {
		newMongoWorkgroupsData = addMongoWorkgroupData(job, city, newMongoWorkgroupsData, *mongoProjectData)
	}
	mongoProjectData.MongoWorkgroupsData[city] = newMongoWorkgroupsData
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

func UpdateMongoProjectData(
	mongoProjectData *internal.MongoProjectData, localProjectData libJenkins.Project, city string, userLogs libJenkins.JenkinsCredentials,
	projectCollection *mongo.Collection,
) error {
	if time.Since(time.Unix(mongoProjectData.LastUpdate, 0)).Hours() < float64(12) {
		err := resetWorkgroupsRemainingBuilds(mongoProjectData, city, projectCollection)
		if err != nil {
			return errors.New(UpdateMongoProjectDataError + err.Error())
		}
		return nil
	}

	jobs, err := libJenkins.GetJobsByProject(localProjectData.Job, city, userLogs)
	if err != nil {
		return errors.New(UpdateMongoProjectDataError + err.Error())
	}
	err = updateMongoWorkgroupsData(mongoProjectData, jobs, city, projectCollection)
	if err != nil {
		return errors.New(UpdateMongoProjectDataError + err.Error())
	}
	return nil
}
