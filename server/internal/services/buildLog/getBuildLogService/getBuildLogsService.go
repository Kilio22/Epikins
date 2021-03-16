package getBuildLogService

import (
	"context"
	"log"
	"net/http"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
)

type BuildLogInfo struct {
	BuildLog  []internal.BuildLogElem `json:"buildLog"`
	TotalPage int64                   `json:"totalPage"`
}

const GetBuildLogError = "cannot get build log"

func getFilter(project string, starter string, city string) bson.M {
	return bson.M{
		"city": bson.M{
			"$regex": primitive.Regex{
				Pattern: "^" + city, Options: "",
			},
		},
		"project": bson.M{
			"$regex": primitive.Regex{
				Pattern: "^" + project, Options: "",
			},
		},
		"starter": bson.M{
			"$regex": primitive.Regex{
				Pattern: "^" + starter, Options: "",
			},
		},
	}
}

func getPaginatedData(buildLogCollection *mongo.Collection, buildLog *[]internal.BuildLogElem, page int64, filter bson.M) (
	*PaginatedData, internal.MyError,
) {
	var paginatedData *PaginatedData
	var err error

	paginatedData, err = New(buildLogCollection).Context(context.TODO()).Limit(20).Page(page).Sort("time", -1).Filter(filter).Decode(buildLog).Find()
	if err != nil {
		log.Println(err)
		return nil, util.GetMyError(GetBuildLogError, err, http.StatusInternalServerError)
	}
	return paginatedData, internal.MyError{}
}

func GetBuildLogService(page int64, project string, starter string, city string, appData *internal.AppData) (
	BuildLogInfo, internal.MyError,
) {
	var buildLog []internal.BuildLogElem
	paginatedData, myError := getPaginatedData(appData.BuildLogCollection, &buildLog, page, getFilter(project, starter, city))
	if myError.Message != "" {
		log.Println(myError)
		return BuildLogInfo{}, myError
	}
	return BuildLogInfo{
		TotalPage: paginatedData.Pagination.TotalPage,
		BuildLog:  buildLog,
	}, internal.MyError{}
}
