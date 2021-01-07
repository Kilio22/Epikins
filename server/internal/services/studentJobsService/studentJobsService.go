package studentJobsService

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"epikins-api/config"
	"epikins-api/internal"
	"epikins-api/pkg/libJenkins"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Module struct {
	ScholarYear string `json:"scolaryear"`
	CodeModule  string `json:"codemodule"`
}

type UserInformationIntraResponse struct {
	Modules []Module `json:"modules"`
}

func getMyHttpError(err error, code int) internal.MyError {
	return internal.MyError{
		Err:        errors.New("cannot get student info on Epitech intranet: " + err.Error()),
		StatusCode: code,
	}
}

func getModulesFromIntraResponse(res *http.Response) ([]Module, internal.MyError) {
	var intraResponse UserInformationIntraResponse
	err := json.NewDecoder(res.Body).Decode(&intraResponse)
	_ = res.Body.Close()
	if err != nil {
		log.Println(err)
		return []Module{}, getMyHttpError(err, http.StatusInternalServerError)
	}
	return intraResponse.Modules, internal.MyError{}
}

func getStudentRegisteredModules(studentEmail string) ([]Module, internal.MyError) {
	req, err := http.NewRequest(http.MethodGet, config.IntraAutologinLink+"/user/"+studentEmail+"/notes?format=json", nil)
	if err != nil {
		log.Println(err)
		return []Module{}, getMyHttpError(err, http.StatusInternalServerError)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []Module{}, getMyHttpError(err, http.StatusInternalServerError)
	}
	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusIMUsed {
		return []Module{}, getMyHttpError(errors.New("bad response code when making request to Epitech intranet, got: "+strconv.Itoa(res.StatusCode)), http.StatusInternalServerError)
	}
	return getModulesFromIntraResponse(res)
}

func getModulesNameFromModules(modules []Module) []string {
	var moduleNameList []string

	for _, module := range modules {
		moduleNameList = append(moduleNameList, module.CodeModule)
	}
	return moduleNameList
}

func getMongoProjectsDataFromStudentModules(modules []Module, projectCollection *mongo.Collection) ([]internal.MongoProjectData, error) {
	cursor, err := projectCollection.Find(context.TODO(), bson.M{"module": bson.M{"$in": getModulesNameFromModules(modules)}})
	if err != nil {
		return []internal.MongoProjectData{}, err
	}
	var mongoProjectsData []internal.MongoProjectData
	err = cursor.All(context.TODO(), &mongoProjectsData)
	return mongoProjectsData, err
}

func StudentJobsService(studentEmail string, userLogs libJenkins.JenkinsCredentials, appData *internal.AppData) internal.MyError {
	modules, myError := getStudentRegisteredModules(studentEmail)
	if myError.Err != nil {
		return myError
	}
	mongoProjectsData, err := getMongoProjectsDataFromStudentModules(modules, appData.ProjectsCollection)
}
