package internal

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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

type JenkinsCredentials struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `json:"username" validate:"required"`
	ApiKey   string             `json:"apiKey" validate:"required"`
}

type User struct {
	Email            string   `json:"email" validate:"required,email"`
	Roles            []string `json:"roles" validate:"required"`
	JenkinsAccountId string   `json:"jenkinsAccount" validate:"required"`
}

type AppData struct {
	ProjectsCollection           *mongo.Collection
	JenkinsCredentialsCollection *mongo.Collection
	UsersCollection              *mongo.Collection
	ProjectsData                 map[libJenkins.AccountType]ProjectsData
	AppId                        string
}
