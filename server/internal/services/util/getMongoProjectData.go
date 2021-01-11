package util

import (
	"errors"
	"log"

	"epikins-api/internal"
	"epikins-api/internal/services/util/mongoUtil"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/mongo"
)

const GetMongoProjectDataError = "cannot get mongo project data"

func GetMongoProjectData(
	project libJenkins.Project, city string, userLogs libJenkins.JenkinsCredentials, projectCollection *mongo.Collection) (
	internal.MongoProjectData, error,
) {
	mongoProjectData, err := mongoUtil.FetchMongoProjectData(project.Job.Name, projectCollection)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			jobs, err := libJenkins.GetJobsByProject(project.Job, city, userLogs)
			if err != nil {
				return internal.MongoProjectData{}, errors.New(GetMongoProjectDataError + err.Error())
			}
			return mongoUtil.AddMongoProjectData(GetNewMongoProjectData(project, GetMongoWorkgroupsDataFromJobs(jobs, city)), projectCollection)
		}
		log.Println(err)
		return internal.MongoProjectData{}, errors.New(GetMongoProjectDataError + ": " + err.Error())
	}
	return mongoProjectData, nil
}
