package studentJobsService

import (
	"net/http"
	"strings"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
	"epikins-api/pkg/libJenkins"
)

type Module struct {
	CodeModule string `json:"codemodule"`
}

type StudentProject struct {
	Module     string `json:"module"`
	Name       string `json:"name"`
	BuildLimit int    `json:"buildLimit"`
}

type StudentJob struct {
	City               string                      `json:"city"`
	MongoWorkgroupData internal.MongoWorkgroupData `json:"mongoWorkgroupData"`
	Project            StudentProject              `json:"project"`
}

const StudentJobsError = "cannot get student jobs"

func getStudentWorkgroup(studentName string, city string, mongoWorkgroupsData internal.MongoProjectData) (
	internal.MongoWorkgroupData, bool,
) {
	for _, mongoWorkgroupData := range mongoWorkgroupsData.MongoWorkgroupsData[city] {
		if strings.Contains(mongoWorkgroupData.Name, studentName) {
			return mongoWorkgroupData, true
		}
	}
	return internal.MongoWorkgroupData{}, false
}

func getStudentJobsFromMongoProjectsData(studentName string, city string, mongoProjectsData []internal.MongoProjectData) []StudentJob {
	var studentJobs []StudentJob

	for _, mongoProjectData := range mongoProjectsData {
		studentWorkgroup, ok := getStudentWorkgroup(studentName, city, mongoProjectData)
		if !ok {
			continue
		}
		studentJobs = append(studentJobs, StudentJob{
			City:               city,
			MongoWorkgroupData: studentWorkgroup,
			Project: StudentProject{
				Module:     mongoProjectData.Module,
				Name:       mongoProjectData.Name,
				BuildLimit: mongoProjectData.BuildLimit,
			},
		})
	}
	return studentJobs
}

func StudentJobsService(studentEmail string, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) (
	[]StudentJob, internal.MyError,
) {
	modules, myError := getStudentRegisteredModules(studentEmail)
	if myError.Message != "" {
		return nil, myError
	}
	if len(modules) == 0 {
		return []StudentJob{}, internal.MyError{}
	}

	city, myError := getStudentCity(studentEmail)
	if myError.Message != "" {
		return nil, myError
	}

	mongoProjectsData, err := getMongoProjectsDataFromStudentModules(modules, city, userLogs, appData)
	if err != nil {
		return nil, util.GetMyError(StudentJobsError, err, http.StatusInternalServerError)
	}
	if len(mongoProjectsData) == 0 {
		return []StudentJob{}, internal.MyError{}
	}

	mongoProjectsData, myError = updateMongoProjectsData(mongoProjectsData, city, userLogs, appData)
	if myError.Message != "" {
		return nil, myError
	}
	return getStudentJobsFromMongoProjectsData(util.GetUsernameFromEmail(studentEmail), city, mongoProjectsData), internal.MyError{}
}
