package internal

import (
	"time"

	"epikins-api/pkg/libJenkins"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectsData struct {
	ProjectList []libJenkins.Job
	LastUpdate  time.Time
}

type MongoWorkgroupData struct {
	Name            string `json:"name"`
	RemainingBuilds int    `json:"remainingBuilds"`
}

type MongoProjectData struct {
	Name                string               `json:"name"`
	MongoWorkgroupsData []MongoWorkgroupData `json:"mongoWorkgroupsData"`
	LastUpdate          int64                `json:"lastUpdate"`
}

type Role string

type User struct {
	Email        string `json:"email" validate:"required,email"`
	Roles        []Role `json:"roles" validate:"required"`
	JenkinsLogin string `json:"jenkinsLogin" validate:"required"`
}

type AppData struct {
	ProjectsCollection           *mongo.Collection
	JenkinsCredentialsCollection *mongo.Collection
	UsersCollection              *mongo.Collection
	ProjectsData                 map[string]ProjectsData
	AppId                        string
}
