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
	BuildLimit          int                  `json:"buildLimit"`
	LastUpdate          int64                `json:"lastUpdate"`
	Module              string               `json:"module"`
	MongoWorkgroupsData []MongoWorkgroupData `bson:"mongoworkgroupsdata,omitempty" json:"mongoWorkgroupsData"`
	Name                string               `json:"name"`
}

// Last update -> la dernière fois que la liste en elle même a été udpate
// Si < 24h -> concidérer la liste comme valable
// Si > 24h -> update la liste : garder les groupes encore présents avec les valeurs actuelles en base, virer ceux qui sont en base mais plus dans la requête
// Comme ça on peut get tout les projets comportant tel ou tel module (requête à l'intra) (add un champ Module string)
// on update si besoin comme précisé ci-dessus, on regarde s'il y a un grp avec l'email donnée
// Si oui on met le projet dans la réponse avec le nombre de build restants

// Add l'url d'un workgroup pour soulager l'api au moment du build, donc faire la requête à jenkins uniquement si le last update est > 24h
// Coté élève, faire une requête à jenkins pour maj les groupes uniquement si le last update > 24

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
