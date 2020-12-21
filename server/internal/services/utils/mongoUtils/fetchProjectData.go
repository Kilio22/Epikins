package mongoUtils

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"

	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
)

func FetchProjectData(project string, jobs []libJenkins.Job, collection *mongo.Collection) (internal.MongoProjectData, error) {
	var projectData internal.MongoProjectData
	err := collection.FindOne(
		context.TODO(), bson.M{"name": project}).Decode(&projectData)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return addProject(project, jobs, collection)
		}
		log.Println(err)
		return internal.MongoProjectData{}, errors.New("cannot fetch data in DB: " + err.Error())
	}

	if shouldResetWorkgroupsRemainingBuilds(projectData) {
		err = resetWorkgroupsRemainingBuilds(&projectData, collection)
		if err != nil {
			return internal.MongoProjectData{}, errors.New("cannot update workgroups remaining builds: " + err.Error())
		}
	}
	return projectData, nil
}
