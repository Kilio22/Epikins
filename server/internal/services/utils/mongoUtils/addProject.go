package mongoUtils

import (
	"context"
	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func addProject(project string, jobs []libJenkins.Job, collection *mongo.Collection) (internal.MongoProjectData, error) {
	var mongoWorkgroupsData []internal.MongoWorkgroupData
	for _, job := range jobs {
		mongoWorkgroupsData = append(mongoWorkgroupsData, internal.MongoWorkgroupData{
			Name:            job.Name,
			RemainingBuilds: config.DefaultBuildNb,
		})
	}
	newProjectData := internal.MongoProjectData{
		Name:                project,
		MongoWorkgroupsData: mongoWorkgroupsData,
		LastUpdate:          getLastMondayDate(),
	}
	_, err := collection.InsertOne(context.TODO(), newProjectData)
	if err != nil {
		log.Println(err)
		return internal.MongoProjectData{}, errors.New("cannot add new project in db: " + err.Error())
	}
	return newProjectData, nil
}
