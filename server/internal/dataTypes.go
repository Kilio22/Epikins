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

type AppData struct {
	Collection   *mongo.Collection
	ProjectsData map[libJenkins.AccountType]ProjectsData
	AppId string
}
