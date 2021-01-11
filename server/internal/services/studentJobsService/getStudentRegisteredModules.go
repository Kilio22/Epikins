package studentJobsService

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
)

type UserInformationIntraResponse struct {
	Modules []Module `json:"modules"`
}

var IntraAutologinLink = util.GetEnvVariable("INTRA_AUTOLOGIN_LINK")

const GetStudentInfoError = "cannot get student info on Epitech intranet"

func getModulesFromIntraResponse(res *http.Response) ([]Module, internal.MyError) {
	var intraResponse UserInformationIntraResponse
	err := json.NewDecoder(res.Body).Decode(&intraResponse)
	_ = res.Body.Close()
	if err != nil {
		log.Println(err)
		return nil, util.GetMyError(GetStudentInfoError, err, http.StatusInternalServerError)
	}
	return intraResponse.Modules, internal.MyError{}
}

func getStudentRegisteredModules(studentEmail string) ([]Module, internal.MyError) {
	req, err := http.NewRequest(http.MethodGet, IntraAutologinLink+"/user/"+studentEmail+"/notes?format=json", nil)
	if err != nil {
		log.Println(err)
		return nil, util.GetMyError(GetStudentInfoError, err, http.StatusInternalServerError)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, util.GetMyError(GetStudentInfoError, err, http.StatusInternalServerError)
	}
	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusIMUsed {
		return nil, util.GetMyError(GetStudentInfoError,
			errors.New("bad response code when making request to Epitech intranet, got: "+strconv.Itoa(res.StatusCode)),
			http.StatusInternalServerError)
	}
	return getModulesFromIntraResponse(res)
}
