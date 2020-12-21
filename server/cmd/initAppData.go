package main

import (
	"epikins-api/internal"
	"epikins-api/internal/services/utils"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/mongo"
)

const ProjectsCollectionName string = "projects"
const JenkinsCredentialsCollectionName string = "credentials"
const UsersCollectionName string = "users"

func initAppData(mongoClient *mongo.Client) *internal.AppData {
	epikinsDatabase := mongoClient.Database(utils.GetEnvVariable("MONGO_DB"))
	projectsCollection := epikinsDatabase.Collection(ProjectsCollectionName)
	credentialsCollection := epikinsDatabase.Collection(JenkinsCredentialsCollectionName)
	usersCollection := epikinsDatabase.Collection(UsersCollectionName)

	return &internal.AppData{
		ProjectsCollection:           projectsCollection,
		JenkinsCredentialsCollection: credentialsCollection,
		UsersCollection:              usersCollection,
		ProjectsData:                 make(map[libJenkins.AccountType]internal.ProjectsData),
		AppId:                        utils.GetEnvVariable("APP_ID"),
	}
}
