package controllerUtil

import (
	"context"

	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetJenkinsCredentials(login string, jenkinsCredentialsCollection *mongo.Collection) (libJenkins.JenkinsCredentials, error) {
	var associatedJenkinsCredentials libJenkins.JenkinsCredentials

	err := jenkinsCredentialsCollection.FindOne(context.TODO(), bson.M{"login": login}).Decode(&associatedJenkinsCredentials)
	if err != nil {
		return libJenkins.JenkinsCredentials{}, err
	}
	return associatedJenkinsCredentials, nil
}
