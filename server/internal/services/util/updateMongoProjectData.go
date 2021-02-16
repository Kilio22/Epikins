package util

import (
	"errors"
	"time"

	"epikins-api/config"
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
			mongoWorkgroupData.LastBuildReset = mongoUtil.GetLastMondayDate()
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
	if mongoProjectData.MongoWorkgroupsData == nil {
		mongoProjectData.MongoWorkgroupsData = map[string][]internal.MongoWorkgroupData{}
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
	if time.Since(time.Unix(mongoProjectData.LastUpdate, 0)).Hours() < config.ProjectJobsRefreshTime && len(mongoProjectData.MongoWorkgroupsData[city]) != 0 {
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
	if len(jobs) != 0 {
		err = updateMongoWorkgroupsData(mongoProjectData, jobs, city, projectCollection)
		if err != nil {
			return errors.New(UpdateMongoProjectDataError + err.Error())
		}
		return nil
	}
	err = mongoUtil.DeleteMongoProjectData(mongoProjectData.Name, projectCollection)
	if err != nil {
		return errors.New(UpdateMongoProjectDataError + err.Error())
	}
	return errors.New(UpdateMongoProjectDataError + "project does not exists on jenkins")
}
