package main

import (
	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/mongo"
)

const BuildLogCollectionName string = "buildLog"
const ProjectsCollectionName string = "projects"
const JenkinsCredentialsCollectionName string = "credentials"
const UsersCollectionName string = "users"

func initAppData(mongoClient *mongo.Client) *internal.AppData {
	epikinsDatabase := mongoClient.Database(util.GetEnvVariable(config.MongoDbKey))
	buildLogCollection := epikinsDatabase.Collection(BuildLogCollectionName)
	projectsCollection := epikinsDatabase.Collection(ProjectsCollectionName)
	credentialsCollection := epikinsDatabase.Collection(JenkinsCredentialsCollectionName)
	usersCollection := epikinsDatabase.Collection(UsersCollectionName)

	return &internal.AppData{
		AppId:                        util.GetEnvVariable(config.AppIdKey),
		BuildLogCollection:           buildLogCollection,
		JenkinsCredentialsCollection: credentialsCollection,
		ProjectsCollection:           projectsCollection,
		ProjectsData:                 make(map[string]internal.ProjectsData),
		UsersCollection:              usersCollection,
	}
}
