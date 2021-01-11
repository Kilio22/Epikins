package studentJobsService

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"epikins-api/internal"
	"epikins-api/internal/services/util"
)

type UserCityIntraResponse struct {
	Location string `json:"location"`
}

func getCityFromIntraResponse(res *http.Response) (string, internal.MyError) {
	var intraResponse UserCityIntraResponse
	err := json.NewDecoder(res.Body).Decode(&intraResponse)
	_ = res.Body.Close()
	if err != nil {
		log.Println(err)
		return "", util.GetMyError(GetStudentInfoError, err, http.StatusInternalServerError)
	}
	if intraResponse.Location == "" {
		return "", util.GetMyError(GetStudentInfoError, errors.New("couldn't get city associated to connected user"), http.StatusInternalServerError)
	}
	splittedLocation := strings.Split(intraResponse.Location, "/")
	if len(splittedLocation) != 2 {
		return "", util.GetMyError(GetStudentInfoError, errors.New("couldn't get city associated to connected user"), http.StatusInternalServerError)
	}
	return splittedLocation[1], internal.MyError{}
}

func getStudentCity(studentEmail string) (string, internal.MyError) {
	req, err := http.NewRequest(http.MethodGet, IntraAutologinLink+"/user/"+studentEmail+"?format=json", nil)
	if err != nil {
		log.Println(err)
		return "", util.GetMyError(GetStudentInfoError, err, http.StatusInternalServerError)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", util.GetMyError("cannot get student info on Epitech intranet", err, http.StatusInternalServerError)
	}
	if res.StatusCode < http.StatusOK || res.StatusCode > http.StatusIMUsed {
		return "", util.GetMyError(GetStudentInfoError,
			errors.New("bad response code when making request to Epitech intranet, got: "+strconv.Itoa(res.StatusCode)),
			http.StatusInternalServerError)
	}
	return getCityFromIntraResponse(res)
}
