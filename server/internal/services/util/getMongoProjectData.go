package util

import (
	"errors"
	"log"

	"epikins-api/internal"
	"epikins-api/internal/services/util/mongoUtil"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetMongoProjectData(
	project libJenkins.Project, jobs []libJenkins.Job, projectCollection *mongo.Collection) (internal.MongoProjectData, error) {
	mongoProjectData, err := mongoUtil.FetchMongoProjectData(project.Job.Name, projectCollection)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return mongoUtil.AddMongoProjectData(project, jobs, projectCollection)
		}
		log.Println(err)
		return internal.MongoProjectData{}, errors.New("cannot fetch data in DB: " + err.Error())
	}
	return mongoProjectData, nil
}
