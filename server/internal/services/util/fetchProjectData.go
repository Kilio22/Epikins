package util

import (
	"epikins-api/internal"
	"epikins-api/internal/services/util/mongoUtil"
	"epikins-api/pkg/libJenkins"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func FetchProjectData(project string, jobs []libJenkins.Job, collection *mongo.Collection) (internal.MongoProjectData, error) {
	mongoProjectData, err := mongoUtil.FetchMongoProjectData(project, collection)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return mongoUtil.AddMongoProjectData(project, jobs, collection)
		}
		log.Println(err)
		return internal.MongoProjectData{}, errors.New("cannot fetch data in DB: " + err.Error())
	}

	if shouldResetWorkgroupsRemainingBuilds(mongoProjectData) {
		err = resetWorkgroupsRemainingBuilds(&mongoProjectData, collection)
		if err != nil {
			return internal.MongoProjectData{}, errors.New("cannot update workgroups remaining builds: " + err.Error())
		}
	}
	return mongoProjectData, nil
}
