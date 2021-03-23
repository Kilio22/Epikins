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
	if mongoWorkgroupData, ok := HasMongoWorkgroupData(job.Name, mongoProjectData.CitiesData[city].MongoWorkgroupsData); ok {
		if shouldResetWorkgroupRemainingBuilds(mongoWorkgroupData) {
			mongoWorkgroupData.RemainingBuilds = mongoProjectData.BuildLimit
			mongoWorkgroupData.LastBuildReset = mongoUtil.GetLastMondayDate()
		}
		mongoWorkgroupData.Url = job.Url
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
	if mongoProjectData.CitiesData == nil {
		mongoProjectData.CitiesData = map[string]internal.CityData{}
	}
	mongoProjectData.CitiesData[city] = internal.CityData{
		LastUpdate:          time.Now().Unix(),
		MongoWorkgroupsData: newMongoWorkgroupsData,
	}
	return mongoUtil.UpdateProject(mongoProjectData.Name, mongoProjectData.Module,
		bson.M{
			"$set": bson.M{
				"citiesdata": mongoProjectData.CitiesData,
			},
		},
		projectCollection,
	)
}

func UpdateMongoProjectData(
	mongoProjectData *internal.MongoProjectData, localProjectData libJenkins.Project, city string, forceUpdate bool,
	userLogs libJenkins.JenkinsCredentials,
	projectCollection *mongo.Collection,
) error {
	if time.Since(time.Unix(mongoProjectData.CitiesData[city].LastUpdate, 0)).Hours() < config.ProjectJobsRefreshTime && len(mongoProjectData.CitiesData[city].MongoWorkgroupsData) != 0 && !forceUpdate {
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
	err = mongoUtil.DeleteMongoProjectData(mongoProjectData.Name, mongoProjectData.Module, projectCollection)
	if err != nil {
		return errors.New(UpdateMongoProjectDataError + err.Error())
	}
	return errors.New(UpdateMongoProjectDataError + "project does not exists on jenkins")
}
