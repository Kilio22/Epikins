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

type UserInformationIntraResponse struct {
	Modules []Module `json:"modules"`
}

type StudentJob struct {
	MongoWorkgroupData internal.MongoWorkgroupData `json:"mongoWorkgroupData"`
	Project            string                      `json:"project"`
}

const StudentJobsError = "cannot get student jobs"

func getStudentWorkgroup(studentName string, mongoWorkgroupsData internal.MongoProjectData) (internal.MongoWorkgroupData, bool) {
	for _, mongoWorkgroupData := range mongoWorkgroupsData.MongoWorkgroupsData {
		if strings.Contains(mongoWorkgroupData.Name, studentName) {
			return mongoWorkgroupData, true
		}
	}
	return internal.MongoWorkgroupData{}, false
}

func getStudentNameFromEmail(studentEmail string) string {
	return strings.Split(studentEmail, "@")[0]
}

func getStudentJobsFromMongoProjectsData(studentName string, mongoProjectsData []internal.MongoProjectData) []StudentJob {
	var studentJobs []StudentJob

	for _, mongoProjectData := range mongoProjectsData {
		studentWorkgroup, ok := getStudentWorkgroup(studentName, mongoProjectData)
		if !ok {
			continue
		}
		studentJobs = append(studentJobs, StudentJob{
			MongoWorkgroupData: studentWorkgroup,
			Project:            mongoProjectData.Name,
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

	mongoProjectsData, err := getMongoProjectsDataFromStudentModules(modules, userLogs, appData)
	if err != nil {
		return nil, util.GetMyError(StudentJobsError, err, http.StatusInternalServerError)
	}
	if len(mongoProjectsData) == 0 {
		return []StudentJob{}, internal.MyError{}
	}

	mongoProjectsData, myError = updateMongoProjectsData(mongoProjectsData, userLogs, appData)
	if myError.Message != "" {
		return nil, myError
	}
	return getStudentJobsFromMongoProjectsData(getStudentNameFromEmail(studentEmail), mongoProjectsData), internal.MyError{}
}
