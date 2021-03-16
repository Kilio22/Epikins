package exportBuildLogService

import (
	"bytes"
	"context"
	"encoding/csv"
	"net/http"
	"strconv"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const ExtractBuildLogError = "cannot retrieve build log"

func getBuildLog(start int64, end int64, project string, city string, buildLogCollection *mongo.Collection) (
	[]internal.BuildLogElem, error,
) {
	cursor, err := buildLogCollection.Find(context.TODO(), bson.M{
		"time": bson.M{"$gt": start, "$lt": end}, "project": bson.M{
			"$regex": primitive.Regex{
				Pattern: "^" + project, Options: "",
			},
		},
		"city": bson.M{
			"$regex": primitive.Regex{
				Pattern: "^" + city, Options: "",
			},
		},
	})
	if err != nil {
		return nil, err
	}

	var buildLog []internal.BuildLogElem
	err = cursor.All(context.TODO(), &buildLog)
	if err != nil {
		return nil, err
	}
	return buildLog, nil
}

func convertBuildLog(buildLog []internal.BuildLogElem) [][]string {
	convertedBuildLog := [][]string{
		{"module", "project", "starter", "target", "time"},
	}
	for _, log := range buildLog {
		timestampString := strconv.FormatInt(log.Time, 10)
		convertedBuildLog = append(convertedBuildLog, []string{
			log.Module,
			log.Project,
			log.Starter,
			log.Target,
			timestampString,
		})
	}
	return convertedBuildLog
}

func ExportBuildLogService(start int64, end int64, project string, city string, appData *internal.AppData) ([]byte, internal.MyError) {
	buildLog, err := getBuildLog(start, end, project, city, appData.BuildLogCollection)
	if err != nil {
		return nil, util.GetMyError(ExtractBuildLogError, err, http.StatusInternalServerError)
	}

	convertedBuildLog := convertBuildLog(buildLog)
	buff := bytes.Buffer{}

	err = csv.NewWriter(&buff).WriteAll(convertedBuildLog)
	if err != nil {
		return nil, util.GetMyError(ExtractBuildLogError, err, http.StatusInternalServerError)
	}
	return buff.Bytes(), internal.MyError{}
}
