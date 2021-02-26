package main

import (
	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/mongo"
)

const ProjectsCollectionName string = "projects"
const JenkinsCredentialsCollectionName string = "credentials"
const UsersCollectionName string = "users"

func initAppData(mongoClient *mongo.Client) *internal.AppData {
	epikinsDatabase := mongoClient.Database(util.GetEnvVariable("MONGO_INITDB_DATABASE"))
	projectsCollection := epikinsDatabase.Collection(ProjectsCollectionName)
	credentialsCollection := epikinsDatabase.Collection(JenkinsCredentialsCollectionName)
	usersCollection := epikinsDatabase.Collection(UsersCollectionName)

	return &internal.AppData{
		ProjectsCollection:           projectsCollection,
		JenkinsCredentialsCollection: credentialsCollection,
		UsersCollection:              usersCollection,
		ProjectsData:                 make(map[string]internal.ProjectsData),
		AppId:                        util.GetEnvVariable("APP_ID"),
	}
}
