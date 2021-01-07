package mongoUtil

import (
	"context"
	"errors"
	"log"
	"time"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddMongoProjectData(project libJenkins.Project, jobs []libJenkins.Job, collection *mongo.Collection) (
	internal.MongoProjectData, error,
) {
	var mongoWorkgroupsData []internal.MongoWorkgroupData
	for _, job := range jobs {
		mongoWorkgroupsData = append(mongoWorkgroupsData, internal.MongoWorkgroupData{
			Url:             job.Url,
			Name:            job.Name,
			RemainingBuilds: config.DefaultBuildNb,
			LastBuildReset:  GetLastMondayDate(),
		})
	}
	newProjectData := internal.MongoProjectData{
		BuildLimit:          config.DefaultBuildNb,
		LastUpdate:          time.Now().Unix(),
		Module:              project.Module,
		MongoWorkgroupsData: mongoWorkgroupsData,
		Name:                project.Job.Name,
	}
	_, err := collection.InsertOne(context.TODO(), newProjectData)
	if err != nil {
		log.Println(err)
		return internal.MongoProjectData{}, errors.New("cannot add new projectName in db: " + err.Error())
	}
	return newProjectData, nil
}
