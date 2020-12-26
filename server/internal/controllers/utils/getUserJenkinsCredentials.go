package utils

import (
	"context"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUserJenkinsCredentials(email string, userCollection *mongo.Collection, jenkinsCredentialsCollection *mongo.Collection) (libJenkins.JenkinsCredentials, error) {
	var user internal.User
	err := userCollection.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		return libJenkins.JenkinsCredentials{}, err
	}

	var associatedJenkinsCredentials libJenkins.JenkinsCredentials
	err = jenkinsCredentialsCollection.FindOne(context.TODO(), bson.M{"login": user.JenkinsLogin}).Decode(&associatedJenkinsCredentials)
	if err != nil {
		return libJenkins.JenkinsCredentials{}, err
	}
	return associatedJenkinsCredentials, nil
}
