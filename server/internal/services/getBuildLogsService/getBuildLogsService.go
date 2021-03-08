package getBuildLogsService

import (
	"context"
	"log"
	"net/http"

	. "github.com/gobeam/mongo-go-pagination"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/bson"
)

type BuildLogsInfo struct {
	BuildLogs []internal.BuildLog `json:"buildLogs"`
	TotalPage int64               `json:"totalPage"`
}

const GetBuildLogsError = "cannot get build log"

func GetBuildLogsService(page int64, appData *internal.AppData) (BuildLogsInfo, internal.MyError) {
	aggPaginatedData, err := New(appData.BuildLogCollection).Context(context.TODO()).Limit(30).Page(page).Sort("time", -1).Aggregate(bson.M{})
	if err != nil {
		log.Println(err)
		return BuildLogsInfo{}, util.GetMyError(GetBuildLogsError, err, http.StatusInternalServerError)
	}

	var buildLogs []internal.BuildLog
	for _, raw := range aggPaginatedData.Data {
		var buildLog internal.BuildLog
		if marshallErr := bson.Unmarshal(raw, &buildLog); marshallErr == nil {
			buildLogs = append(buildLogs, buildLog)
		}
	}
	return BuildLogsInfo{
		TotalPage: aggPaginatedData.Pagination.TotalPage,
		BuildLogs: buildLogs,
	}, internal.MyError{}
}
