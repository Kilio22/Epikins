package internal

import (
	"time"

	"epikins-api/pkg/libJenkins"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectsData struct {
	LastUpdate  time.Time
	ProjectList []libJenkins.Project
}

type MongoWorkgroupData struct {
	Name            string `json:"name"`
	LastBuildReset  int64  `json:"lastBuildReset"`
	RemainingBuilds int    `json:"remainingBuilds"`
	Url             string `json:"url"`
}

type CityData struct {
	LastUpdate          int64                `json:"lastUpdate"`
	MongoWorkgroupsData []MongoWorkgroupData `bson:"mongoworkgroupsdata,omitempty" json:"mongoWorkgroupsData"`
}

type MongoProjectData struct {
	BuildLimit int                 `json:"buildLimit"`
	CitiesData map[string]CityData `bson:"citiesdata,omitempty" json:"citiesData"`
	Module     string              `json:"module"`
	Name       string              `json:"name"`
}

type Role string

type User struct {
	Email        string `json:"email" validate:"required,email"`
	JenkinsLogin string `json:"jenkinsLogin" validate:"required"`
	Roles        []Role `json:"roles" validate:"required"`
}

type BuildLogElem struct {
	City    string `json:"city"`
	Module  string `json:"module"`
	Project string `json:"project"`
	Starter string `json:"starter"`
	Target  string `json:"target"`
	Time    int64  `json:"time"`
}

type AppData struct {
	AppId                        string
	BuildLogCollection           *mongo.Collection
	JenkinsCredentialsCollection *mongo.Collection
	ProjectsCollection           *mongo.Collection
	ProjectsData                 map[string]ProjectsData
	UsersCollection              *mongo.Collection
}
