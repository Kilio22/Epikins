package internal

import (
	"time"

	"epikins-api/pkg/libJenkins"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectsData struct {
	ProjectList []libJenkins.Project
	LastUpdate  time.Time
}

type MongoWorkgroupData struct {
	Name            string `json:"name"`
	LastBuildReset  int64  `json:"lastBuildReset"`
	RemainingBuilds int    `json:"remainingBuilds"`
	Url             string `json:"url"`
}

type MongoProjectData struct {
	BuildLimit          int                             `json:"buildLimit"`
	LastUpdate          int64                           `json:"lastUpdate"`
	Module              string                          `json:"module"`
	MongoWorkgroupsData map[string][]MongoWorkgroupData `bson:"mongoworkgroupsdata,omitempty" json:"mongoWorkgroupsData"`
	Name                string                          `json:"name"`
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
