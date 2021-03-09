package getBuildLogsService

import (
	"context"
	"log"
	"net/http"
	"time"

	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
)

type BuildLogsInfo struct {
	BuildLogs []internal.BuildLogs `json:"buildLogs"`
	TotalPage int64                `json:"totalPage"`
}

const GetBuildLogsError = "cannot get build logs"

func getFilter(project string, starter string) bson.M {
	return bson.M{
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

func getPaginatedData(buildLogCollection *mongo.Collection, buildLogs *[]internal.BuildLogs, page int64, project string, starter string) (
	*PaginatedData, internal.MyError,
) {
	var paginatedData *PaginatedData
	var err error

	paginatedData, err = New(buildLogCollection).Context(context.TODO()).Limit(20).Page(page).Sort("time", -1).Filter(getFilter(project, starter)).Decode(buildLogs).Find()
	if err != nil {
		log.Println(err)
		return nil, util.GetMyError(GetBuildLogsError, err, http.StatusInternalServerError)
	}
	return paginatedData, internal.MyError{}
}

func GetBuildLogsService(page int64, project string, starter string, appData *internal.AppData) (BuildLogsInfo, internal.MyError) {
	buildLogs := []internal.BuildLogs{}
	paginatedData, myError := getPaginatedData(appData.BuildLogsCollection, &buildLogs, page, project, starter)
	if myError.Message != "" {
		log.Println(myError)
		return BuildLogsInfo{}, myError
	}

	if time.Since(time.Unix(appData.LastBuildLogsCleanup, 0)).Hours() >= 24 {
		currentUnixTimestamp := time.Now().Unix()
		_, _ = appData.BuildLogsCollection.DeleteMany(context.TODO(), bson.M{
			"time": bson.M{"$lte": time.Unix(currentUnixTimestamp, 0).AddDate(0, 0, -30).Unix()},
		})
		appData.LastBuildLogsCleanup = currentUnixTimestamp
	}
	return BuildLogsInfo{
		TotalPage: paginatedData.Pagination.TotalPage,
		BuildLogs: buildLogs,
	}, internal.MyError{}
}
