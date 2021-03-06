package libJenkins

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type JobList struct {
	Jobs []Job
}

func GetJobsByURL(url string, userLogs JenkinsCredentials) ([]Job, error) {
	res, err := makeHttpRequest(http.MethodGet, url, JenkinsAPIPart, userLogs, "")
	if res != nil && res.StatusCode == 404 {
		return []Job{}, nil
	}
	if err != nil {
		return []Job{}, errors.New("cannot get jobs: " + err.Error())
	}

	var jobsList JobList
	err = json.NewDecoder(res.Body).Decode(&jobsList)
	_ = res.Body.Close()
	if err != nil {
		log.Println(err)
		return []Job{}, errors.New("cannot get jobs: " + err.Error())
	}
	return jobsList.Jobs, nil
}
